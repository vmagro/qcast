//
//  ViewController.swift
//  QCast
//
//  Created by Vinnie Magro on 11/9/17.
//  Copyright © 2017 Vinnie Magro. All rights reserved.
//

import UIKit
import Material
import CoreLocation
import Libqcast

class HomeViewController: UIViewController {
  fileprivate var menuButton: IconButton!
  fileprivate var searchButton: IconButton!

  fileprivate var qcastTitle: UILabel!
  fileprivate var nearbyCard: View!
  fileprivate var partiesTable: TableView!
  fileprivate var fabButton: FABButton!
  
  fileprivate var locationManager: CLLocationManager!
  
  fileprivate var loginController: LoginViewController?
  
  fileprivate var partiesGeo: MobilePartiesGeoServiceProtocol?
  fileprivate var parties: [LibqcastParty] = []
  
  var joinAfterCreate: LibqcastParty?
  fileprivate var waitingToJoin: LibqcastParty?

  open override func viewDidLoad() {
    super.viewDidLoad()
    view.backgroundColor = Color.grey.lighten5

    prepareMenuButton()
    prepareNavigationItem()
    preparePartiesList()
    prepareFABButton()
    prepareQCastTitle()
    prepareBackgroundImage()
  }

  open override func viewWillAppear(_ animated: Bool) {
    super.viewWillAppear(animated)

    prepareNavigationItem()
  }
  
  open override func viewDidAppear(_ animated: Bool) {
    super.viewDidAppear(animated)
    
    locationManager = CLLocationManager()
    locationManager.delegate = self
    
    if CLLocationManager.authorizationStatus() == .notDetermined {
      locationManager.requestWhenInUseAuthorization()
    }
    
    locationManager.startUpdatingLocation()
    
    // check if we were able to reach the database and if not show an error
    if !LibqcastCanConnectToDB() {
      let alert = UIAlertController(title: "Connection Error", message: "Unable to connect to QCast, please try again later or check if there's an update on the App Store.", preferredStyle: .alert)
      self.present(alert, animated: true, completion: nil)
    }
    
    if joinAfterCreate != nil {
      joinParty(party: joinAfterCreate!)
      joinAfterCreate = nil
    }
    
    partiesGeo = MobileNewPartiesGeo()
  }
}

extension HomeViewController: LoginViewControllerDelegate {
  func loginDidSucceed() {
    loginController?.dismiss(animated: true, completion: {
      // if we were trying to join a party join it now that we're logged in
      if self.waitingToJoin != nil {
        self.joinParty(party: self.waitingToJoin)
        self.waitingToJoin = nil
      }
    })
  }
  
  func loginCanceled() {
    self.waitingToJoin = nil
    loginController?.dismiss(animated: true, completion: {})
  }
  
  func joinParty(party: LibqcastParty!) {
    // if we aren't logged in yet, show the login view controller
    if !LibqcastAuthService().loggedIn() {
      loginController = LoginViewController()
      loginController?.delegate = self
      loginController?.instructions = "Login with Spotify so you can search for tracks to add to the queue!"
      present(loginController!, animated: true)
      
      // save that we were trying to join this party
      waitingToJoin = party
      return
    }
    
    // join the party, wait for it to be joined and then show the queue view controller
    let partyService = LibqcastRethinkPartyService()
    
    do {
      try partyService?.join(party)
    } catch let error as NSError {
      var text = "Unable to join party.\n Please try again later."
      if error.localizedDescription == "api-version-fail" {
        text = "This party is incompatible with your version of QCast.\n Please check the App Store for updates."
      }
      let alert = UIAlertController(title: "Connection Error", message: text, preferredStyle: .alert)
      alert.addAction(UIAlertAction(title: "Close", style: .default, handler: { (_ action) in
        alert.dismiss(animated: true, completion: nil)
      }))
      self.present(alert, animated: true, completion: nil)
      return
    }
    
    // now that we've joined the party, show the queue
    navigationController?.pushViewController(QueueViewController(rootViewController: QueueRootViewController()), animated: true)
  }
}

fileprivate extension HomeViewController {
  func prepareMenuButton() {
    menuButton = IconButton(image: Icon.cm.settings)
    menuButton.tintColor = Color.white
  }

  func prepareNavigationItem() {
    // make navigation bar clear
    navigationController?.navigationBar.dividerColor = Color.clear
    navigationController?.navigationBar.backgroundColor = Color.clear
    navigationItem.leftViews = [menuButton]
  }

  func prepareQCastTitle() {
    qcastTitle = UILabel()
    qcastTitle.text = "QCast"
    qcastTitle.textColor = Color.white
    qcastTitle.textAlignment = .center
    //        qcastTitle.font = UIFont.systemFont(ofSize: 48, weight: .light)
    qcastTitle.font = RobotoFont.light(with: 48)
    view.layout(qcastTitle).centerHorizontally().top(16)
  }

  func preparePartiesList() {
    nearbyCard = View()

    let toolbar = Toolbar()
    toolbar.title = "Nearby Parties"
    toolbar.titleLabel.textAlignment = .left
//    nearbyCard.toolbar = toolbar
//    nearbyCard.toolbarEdgeInsets.left = 8

    nearbyCard.depthPreset = .depth3
//    nearbyCard.cornerRadius = 4
    nearbyCard.cornerRadiusPreset = .cornerRadius3

    partiesTable = TableView()
    partiesTable.dataSource = self
    partiesTable.delegate = self
    partiesTable.backgroundColor = Color.clear
    partiesTable.showsHorizontalScrollIndicator = false

//    nearbyCard.contentView = partiesTable
    view.layout(nearbyCard).top(128).left(16).right(16).bottom(128)
    nearbyCard.layout(partiesTable).left(8).right(8).top(8).bottom(8)
  }

  func prepareFABButton() {
    fabButton = FABButton(image: Icon.add, tintColor: Color.white)
    fabButton.backgroundColor = Color.green.base
    fabButton.addTarget(self, action: #selector(handleAddPartyButton), for: .touchUpInside)
    view.layout(fabButton).width(64).height(64).bottom(24).right(24)
  }

  func prepareBackgroundImage() {
    let image = UIImageView()
    image.image = UIImage(named: "Concert")
    image.contentMode = .scaleAspectFill
    image.layer.zPosition = -1
    view.layout(image).top(-64).bottom().left().right()
  }

}

extension HomeViewController: MobilePartiesGeoCallbackProtocol {
  
  @objc
  func handleAddPartyButton() {
    navigationController?.pushViewController(NamePartyViewController(), animated: true)
  }

  func partyEntered(_ party: LibqcastParty!) {
    // if we already have this party in the list don't add it again
    for p in self.parties {
      if p.id_() == party.id_() {
        return
      }
    }
    
    parties.append(party)
    DispatchQueue.main.async {
      self.partiesTable.reloadData()
    }
  }
  
  func partyExited(_ party: LibqcastParty!) {
    self.parties = parties.filter({ (p) -> Bool in
      return p.id_() != party.id_()
    })
    DispatchQueue.main.async {
      self.partiesTable.reloadData()
    }
  }
  
  func partiesReady() {
    // TODO: check if we have no parties and if so show an indicator that there were none found
  }
}

extension HomeViewController: UITableViewDataSource {
  func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
    return parties.count
  }

  func numberOfSections(in tableView: UITableView) -> Int {
    return 1
  }

  func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
    let cell = tableView.dequeueReusableCell(withIdentifier: "TableViewCell", for: indexPath) as! TableViewCell

    let party = parties[indexPath.row]
    cell.textLabel?.text = party.name()
    cell.detailTextLabel?.text = "\(party.code()!) - \(party.host().name() ?? "(no host)")"

    cell.dividerColor = Color.grey.lighten2

    return cell
  }
}

extension HomeViewController: UITableViewDelegate {
  func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
    let party = parties[indexPath.row]
    joinParty(party: party)
  }
}

extension HomeViewController: CLLocationManagerDelegate {
  func locationManager(_ manager: CLLocationManager, didUpdateLocations locations: [CLLocation]) {
    // If it's a relatively recent event, turn off updates to save power.
    let location = locations.last!
    let eventDate = location.timestamp
    let howRecent = eventDate.timeIntervalSinceNow
    if (abs(howRecent) < 15.0) {
      // If the event is recent, do something with it.
      NSLog("latitude %+.6f, longitude %+.6f\n",
            location.coordinate.latitude,
            location.coordinate.longitude);
      locationManager.stopUpdatingLocation()
      
      // look for parties nearby
      let loc = LibqcastNewLocation(location.coordinate.latitude, location.coordinate.longitude)
      partiesGeo?.watchParties(loc, callback: self)
      
      do {
//        let service = LibqcastRethinkPartyService()!
//        try service.startParty("Test Party", location: loc, moodPlaylist: nil)
      } catch {
        
      }
    }
  }
  
  func locationManager(_ manager: CLLocationManager, didFailWithError error: Error) {
    NSLog("Location failed %@\n", error as NSError)
  }
  
  func locationManager(_ manager: CLLocationManager, didChangeAuthorization status: CLAuthorizationStatus) {
    if status == .authorizedAlways || status == .authorizedWhenInUse {
      locationManager.startUpdatingLocation()
    }
  }
}
