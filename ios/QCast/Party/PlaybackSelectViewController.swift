////
////  PlaybackSelectViewController
////  QCastDev
////
////  Created by Vinnie Magro on 11/28/17.
////
//
//import Foundation
//import UIKit
//import Material
//import Libqcast
//
//class PlaybackSelectViewController: UIViewController {
//
//  fileprivate var navContainer = UIView()
//  fileprivate var devicesTable = UITableView()
//
//  var party: LibqcastParty!
////  fileprivate var devices: LibqcastPlaybackDeviceList?
//
//  override var preferredStatusBarStyle: UIStatusBarStyle {
//    return .lightContent
//  }
//
//  override func viewDidLoad() {
//    super.viewDidLoad()
//
//    view.backgroundColor = Color.grey.lighten5
//
//    prepareNavContainer()
//    prepareDevicesTable()
//
//    loadDevices(sender: view)
//  }
//
//  func prepareNavContainer() {
//    navContainer.backgroundColor = Color.green.base
//
//    view.layout(navContainer).top(0).left(0).right(0).height(256)
//
//    prepareTitle()
//    prepareDescription()
//    prepareCancelButton()
//    prepareRefreshButton()
//  }
//
//  func prepareCancelButton() {
//    let cancelButton = Button(title: "Cancel", titleColor: Color.white)
//    cancelButton.addTarget(self, action: #selector(self.cancelPressed), for: .touchUpInside)
//    view.layout(cancelButton).top(20).left(14).height(44)
//  }
//
//  func prepareTitle() {
//    let titleLabel = UILabel(frame: CGRect(x: 0, y: 0, width: 200, height: 44))
//    titleLabel.text = "Choose Playback"
//    titleLabel.textColor = Color.white
//    titleLabel.font = UIFont.systemFont(ofSize: 17, weight: .semibold)
//
//    navContainer.layout(titleLabel).centerHorizontally().top(32)
//  }
//
//  func prepareDescription() {
//    let descriptionLabel = UILabel()
//    descriptionLabel.text = "Choose a playback device below to start the music!\n\nYour Spotify Connect devices will appear below, so make sure Spotify is running on your device."
//    descriptionLabel.textColor = Color.white
//    descriptionLabel.font = RobotoFont.light(with: 22)
//    descriptionLabel.backgroundColor = Color.green.base
//    descriptionLabel.numberOfLines = 0
//
//    navContainer.layout(descriptionLabel).left(8).right(8).bottom(16)
//  }
//
//  func prepareDevicesTable() {
//    let tableWrapper = UIView()
//    devicesTable = TableView()
//
//    devicesTable.dataSource = self
//    devicesTable.delegate = self
//    devicesTable.backgroundColor = Color.white
//
//    tableWrapper.cornerRadiusPreset = .cornerRadius3
//    tableWrapper.depthPreset = .depth3
//
//    view.layout(tableWrapper).top(272).left(16).right(16).bottom(64)
//    tableWrapper.layout(devicesTable).left(8).right(8).top(8).bottom(8)
//  }
//
//  func prepareRefreshButton() {
//    let refreshButton = Button(title: "Reload", titleColor: Color.green.base)
//    refreshButton.addTarget(self, action: #selector(self.loadDevices), for: .touchUpInside)
//
//    view.layout(refreshButton).centerHorizontally().bottom(16)
//  }
//
//  @objc
//  func cancelPressed(sender: Button) {
//    self.dismiss(animated: true, completion: nil)
//    //delegate?.loginCanceled()
//  }
//
//  @objc
//  func loadDevices(sender: UIView) {
//    DispatchQueue.global().async {
//      // fire off a request to get playback devices
//      self.devices = try? self.party.player().playbackDevices()
//      print("Found \(self.devices?.numDevices()) devices")
//      DispatchQueue.main.async {
//        self.devicesTable.reloadData()
//      }
//    }
//  }
//
//}
//
//extension PlaybackSelectViewController: UITableViewDataSource {
//  func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
//    if devices == nil {
//      return 0
//    }
//    return devices!.numDevices()
//  }
//
//  func numberOfSections(in tableView: UITableView) -> Int {
//    return 1
//  }
//
//  func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
//    let cell = tableView.dequeueReusableCell(withIdentifier: "TableViewCell", for: indexPath) as! TableViewCell
//
//    let device = try? devices!.device(at: indexPath.row)
//    cell.textLabel?.text = device?.name()
//
//    cell.dividerColor = Color.grey.lighten2
//
//    return cell
//  }
//}
//
//extension PlaybackSelectViewController: UITableViewDelegate {
//  func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
//    // TODO: select the given device
//    let player = LibqcastRethinkPartyService().currentParty().player()
//    let device = try! player!.playbackDevices().device(at: indexPath.row)
//    try? player?.setDevice(device)
//
//    // TODO: check if there was an error?
//    self.dismiss(animated: true, completion: nil)
//  }
//}

