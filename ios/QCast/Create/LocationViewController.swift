//
//  LocationViewController.swift
//  QCastDev
//
//  Created by Vinnie Magro on 11/21/17.
//

import Foundation
import UIKit
import Material
import MapKit
import Libqcast

class LocationViewController: UIViewController {
  
  // in progress party being created
  var partyName: String?
  
  fileprivate var instructionsContainer: UIView!
  fileprivate var map: MKMapView!
  fileprivate var nextButton: Button!
  
  fileprivate var locationManager: CLLocationManager!
  
  open override func viewDidLoad() {
    super.viewDidLoad()
    view.backgroundColor = Color.grey.lighten5
    
    prepareNavigationItem()
    prepareInstructions()
    prepareMap()
    prepareNextButton()
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
  }
}

fileprivate extension LocationViewController {
  
  func prepareNavigationItem() {
    // reset navigation bar from the clear crap in the home screen
    navigationItem.titleLabel.text = "Locate Your Party"
    navigationItem.titleLabel.textColor = Color.white
    
    let b = navigationController?.navigationBar
    b?.backgroundColor = Color.green.base
    b?.tintColor = Color.white
    b?.isTranslucent = false
    b?.dividerColor = Color.clear
    
    navigationItem.backBarButtonItem?.tintColor = Color.white
    navigationItem.backButton.tintColor = Color.white
  }
  
  func prepareInstructions() {
    instructionsContainer = UIView()
    instructionsContainer.backgroundColor = Color.green.base
    
    let label = UILabel()
    label.text = "Tag your party with your location so that your guests can find it more quickly. (optional)"
    label.textColor = Color.white
    label.font = RobotoFont.light(with: 22)
    label.numberOfLines = 0
    
    view.layout(instructionsContainer).left(0).top(0).right(0).height(128)
    instructionsContainer.layout(label).left(8).right(8).bottom(16)
  }
  
  func prepareMap() {
    map = MKMapView()
    
    map.showsUserLocation = true
    map.mapType = .hybrid
    map.isUserInteractionEnabled = false
    
    // zoom in on the current location
//    let noLocation = CLLocationCoordinate2D()
//    let viewRegion = MKCoordinateRegionMakeWithDistance(noLocation, 200, 200)
//    map.setRegion(viewRegion, animated: false)
    
    // TODO: this doesn't seem to work
    map.shapePreset = .square
    view.layout(map).left(8).right(8).top(140).height(200)
  }
  
  func prepareNextButton() {
    nextButton = Button()
    
    nextButton.title = "Start the Party!"
    nextButton.addTarget(self, action: #selector(self.handleNextButton), for: .touchUpInside)
    
    view.layout(nextButton).bottom(14).right(12)
  }
  
}

fileprivate extension LocationViewController {
  @objc
  func handleNextButton() {
    let coord = locationManager.location?.coordinate
    var loc: LibqcastLocation? = nil
    if coord != nil {
      loc = LibqcastNewLocation(coord!.latitude, coord!.longitude)
    }
    do {
      let service = LibqcastRethinkPartyService()!
      let party = try service.startParty(self.partyName, location: loc, moodPlaylist: nil)
      // tell the home view controller to join this party then pop back to it
      let homeVc = navigationController?.viewControllers[0] as? HomeViewController
      homeVc?.joinAfterCreate = party
      navigationController?.popToRootViewController(animated: true)
    } catch {
      let alert = UIAlertController(title: "Oops", message: "Something happened and we weren't able to create your party.\nPlease try again later.", preferredStyle: .alert)
      alert.addAction(UIAlertAction(title: "Close", style: .default, handler: { (_ action) in
        alert.dismiss(animated: true, completion: nil)
      }))
      self.present(alert, animated: true, completion: nil)
    }
  }
  
}

extension LocationViewController: CLLocationManagerDelegate {
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
      
      map.setRegion(MKCoordinateRegion(center: location.coordinate, span: MKCoordinateSpan(latitudeDelta: 0.00625, longitudeDelta: 0.00625)), animated: true)
      
      locationManager.stopUpdatingLocation()
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
