//
//  QueueViewController.swift
//  QCast
//
//  Created by Vinnie Magro on 11/9/17.
//  Copyright © 2017 Vinnie Magro. All rights reserved.
//

import Foundation
import UIKit
import Material
import Dwifft
import Kingfisher

import Libqcast

class QueueRootViewController: UIViewController {

  var tracks: [LibqcastTrack] = [] {
    didSet {
      var ids = [String]()
      for track in tracks {
        ids.append(track.id_())
      }
      
      DispatchQueue.main.async {
        self.differ.rows = ids
      }
    }
  }

  var tableView = TableView()
  var differ: SingleSectionTableViewDiffCalculator<String>!
  fileprivate var controls: PlaybackControlBar?
  
  var queue: MobileQueueProtocol?
  var partyService: LibqcastPartyServiceProtocol!
  
  override init(nibName nibNameOrNil: String?, bundle nibBundleOrNil: Bundle?) {
    super.init(nibName: nibNameOrNil, bundle: nibBundleOrNil)
    
    differ = SingleSectionTableViewDiffCalculator<String>(tableView: tableView, initialRows: [], sectionIndex: 0)
    differ.deletionAnimation = .automatic
    differ.insertionAnimation = .automatic
  }
  
  required init?(coder aDecoder: NSCoder) {
    super.init(coder: aDecoder)
    
    differ = SingleSectionTableViewDiffCalculator<String>(tableView: tableView, initialRows: [], sectionIndex: 0)
    differ.deletionAnimation = .automatic
    differ.insertionAnimation = .automatic
  }

  open override func viewDidLoad() {
    super.viewDidLoad()
    view.backgroundColor = Color.grey.lighten5
    
    // start listening to queue updates
    partyService = LibqcastRethinkPartyService()
    let currentParty = partyService.currentParty()!
    queue = MobileWrapQueue(currentParty.queue())
    
    if currentParty.amIHost() {
      do {
        try currentParty.becomePlayer()
        
        controls?.refresh()
      } catch {
        // TODO: handle error from this function
      }
    }
    
    prepare()
  }

  override func viewDidAppear(_ animated: Bool) {
    super.viewDidAppear(animated)
    // register for queue update
    queue?.registerUpdateListener(self)
    // watch playback events
//    player?.watchEvents(self)
  }

  override func viewDidDisappear(_ animated: Bool) {
    super.viewDidDisappear(animated)
    // unregister for queue update
    queue?.unregisterUpdateListener(self)
    // stop watching playback events
//    player?.stopWatching(self)
  }

  func prepare() {
    prepareTable()
    
    // only enable the control strip when this user is the host
    if partyService.currentParty().amIHost() {
      prepareControls()
    }
  }
  
  func prepareTable() {
    self.tableView.delegate = self
    self.tableView.dataSource = self
    self.tableView.separatorInset = .zero
    self.tableView.showsHorizontalScrollIndicator = false
    
    let bottom = CGFloat(partyService.currentParty().amIHost() ? 64 : 0)
    view.layout(tableView).top(0).left(0).right(0).bottom(bottom)
  }
  
  func prepareControls() {
    controls = PlaybackControlBar()
    controls!.viewController = self
    controls!.playing = true
    controls!.refresh()
    
    SpotifyPlayer.sharedInstance.uiDelegate = controls
    
    view.layout(controls!).bottom(0).height(64).left(0).right(0)
  }

}

extension QueueRootViewController: UITableViewDataSource, UITableViewDelegate {
  func numberOfSections(in tableView: UITableView) -> Int {
    return 1
  }

  func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
    return differ.rows.count
  }

  func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
    let cell = tableView.dequeueReusableCell(withIdentifier: "TableViewCell", for: indexPath) as! TableViewCell

    let trackID = differ.rows[indexPath.row]
    var track: LibqcastTrack!
    for t in self.tracks {
      if t.id_() == trackID {
        track = t
        break
      }
    }

    if indexPath.row == 0 {
      // make a special cell for the current track
      // TODO: use an actual separate cell design cuz the current one looks dumb
      
      cell.textLabel?.text = track.title()
      cell.imageView?.kf.setImage(with: URL(string: track.album().imageURL()), placeholder: Icon.audio?.resize(toHeight: 128)?.resize(toWidth: 128))
      cell.detailTextLabel?.text = track.artistDisplay()

      let avatar = UIImageView(frame: CGRect(x: 0, y: 0, width: 44, height: 44))
      avatar.kf.setImage(with: URL(string: track.addedBy().imageURL()))
      avatar.layer.cornerRadius = avatar.frame.width / 2
      avatar.layer.masksToBounds = true
      cell.accessoryView = avatar
      cell.dividerColor = Color.grey.lighten2

      return cell
    }

    cell.textLabel?.text = track.title()
    cell.imageView?.kf.setImage(with: URL(string: track.album().imageURL()), placeholder: Icon.audio?.resize(toHeight: 64)?.resize(toWidth: 64))
    cell.detailTextLabel?.text = track.artistDisplay()

    let avatar = UIImageView(frame: CGRect(x: 0, y: 0, width: 44, height: 44))
    avatar.kf.setImage(with: URL(string: track.addedBy().imageURL()))
    avatar.layer.cornerRadius = avatar.frame.width / 2
    avatar.layer.masksToBounds = true
    cell.accessoryView = avatar
    cell.dividerColor = Color.grey.lighten2

    cell.layoutMargins = .zero
    cell.separatorInset = .zero
    cell.preservesSuperviewLayoutMargins = false

    return cell
  }

  func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
    if indexPath.row == 0 {
      return 128
    }
    return 64
  }
}

extension QueueRootViewController: MobileQueueCallbackProtocol {
  func queueUpdated(_ queue: MobileQueueProtocol!) {
    self.queue = queue

    var newTracks = [LibqcastTrack]()
    for i in 0..<queue.numTracks() {
      do {
        let track: LibqcastTrack
        try track = queue.track(at: i)
        newTracks.append(track)
      } catch {

      }
    }

    self.tracks = newTracks
    
    controls?.refresh()
  }
}

// FAB stuff

class QueueViewController: FABMenuController {
  fileprivate var menuButton: IconButton!
  fileprivate var searchButton: IconButton!

  fileprivate var fabButton: FABButton!
  fileprivate var playlistsFABMenuItem: FABMenuItem!
  fileprivate var searchFABMenuItem: FABMenuItem!

  open override func viewDidLoad() {
    super.viewDidLoad()
    view.backgroundColor = Color.grey.lighten5

    prepareMenuButton()
    prepareNavigationItem()
    prepareFABButton()
    prepareSearchFABMenuItem()
    preparePlaylistsFABMenuItem()
    prepareFABMenu()
  }

  open override func viewWillAppear(_ animated: Bool) {
    super.viewWillAppear(animated)

    prepareNavigationItem()
  }
}

fileprivate extension QueueViewController {
  func prepareMenuButton() {
    menuButton = IconButton(image: Icon.cm.menu, tintColor: Color.white)
  }

  func prepareNavigationItem() {
    navigationItem.titleLabel.text = LibqcastRethinkPartyService().currentParty().name()
    navigationItem.titleLabel.textColor = Color.white

    let b = navigationController?.navigationBar
    b?.backgroundColor = Color.green.base
    b?.tintColor = Color.white
    b?.isTranslucent = false
    b?.dividerColor = Color.grey.lighten2

    navigationItem.backBarButtonItem?.tintColor = Color.white
    navigationItem.backButton.tintColor = Color.white
  }

  func prepareFABButton() {
    fabButton = FABButton(image: Icon.add, tintColor: Color.white)
    fabButton.backgroundColor = Color.green.base

    view.layout(fabButton).width(64).height(64).bottom(24).right(24)
  }

  fileprivate func prepareSearchFABMenuItem() {
    playlistsFABMenuItem = FABMenuItem()
    playlistsFABMenuItem.title = "Playlists"
    playlistsFABMenuItem.fabButton.image = Icon.cm.audioLibrary
    playlistsFABMenuItem.fabButton.tintColor = .white
    playlistsFABMenuItem.fabButton.pulseColor = .white
    playlistsFABMenuItem.fabButton.backgroundColor = Color.green.base
    playlistsFABMenuItem.fabButton.addTarget(self, action: #selector(handlePlaylistsFAB(button:)), for: .touchUpInside)
  }

  fileprivate func preparePlaylistsFABMenuItem() {
    searchFABMenuItem = FABMenuItem()
    searchFABMenuItem.title = "Search"
    searchFABMenuItem.fabButton.image = Icon.cm.search
    searchFABMenuItem.fabButton.tintColor = .white
    searchFABMenuItem.fabButton.pulseColor = .white
    searchFABMenuItem.fabButton.backgroundColor = Color.green.base
    searchFABMenuItem.fabButton.addTarget(self, action: #selector(handleSearchFAB(button:)), for: .touchUpInside)
  }

  fileprivate func prepareFABMenu() {
    fabMenu.fabButton = fabButton
//    fabMenu.fabMenuItems = [playlistsFABMenuItem, searchFABMenuItem]
    fabMenu.fabMenuItems = [searchFABMenuItem]
    fabMenuBacking = .none

    let bottom = CGFloat(LibqcastRethinkPartyService().currentParty().amIHost() ? 80 : 16)
    view.layout(fabMenu).width(64).height(64).bottom(bottom).right(16)
  }
}

fileprivate extension QueueViewController {

  @objc
  fileprivate func handleSearchFAB(button: UIButton) {
    navigationController?.pushViewController(SearchViewController(), animated: true)
    fabMenu.close()
    fabMenu.fabButton?.animate(.rotate(0))
  }

  @objc
  fileprivate func handlePlaylistsFAB(button: UIButton) {
    //        transition(to: RemindersViewController())
    navigationController?.pushViewController(PlaylistsViewController(), animated: true)
    fabMenu.close()
    fabMenu.fabButton?.animate(.rotate(0))
  }
}

// handle FAB events
fileprivate extension QueueViewController {
  @objc
  func fabMenuWillOpen(fabMenu: FABMenu) {
    fabMenu.fabButton?.animate(.rotate(45))
  }

  @objc
  func fabMenuWillClose(fabMenu: FABMenu) {
    fabMenu.fabButton?.animate(.rotate(0))
  }
}

fileprivate class PlaybackControlBar: Material.View, PlayerUIDelegate {
  
  fileprivate var track = UILabel()
  fileprivate var artist = UILabel()
  fileprivate var playPause: Material.Button!
  fileprivate var skip: Material.Button!

  var playing = true {
    didSet {
      DispatchQueue.main.async {
        if self.playing {
          self.playPause.image = Icon.cm.pause
        } else {
          self.playPause.image = Icon.cm.play
        }
      }
    }
  }

  var viewController: QueueRootViewController!

  override func prepare() {
    super.prepare()

    self.backgroundColor = Material.Color.green.base

    prepareTrack()
    prepareSkip()
    preparePlayPause()
  }
  
  func prepareTrack() {
    track.textColor = Color.white
    track.font = RobotoFont.medium(with: 18)
    artist.textColor = Color.white
    artist.font = RobotoFont.light(with: 16)
    layout(track).left(16).top(8)
    layout(artist).left(16).bottom(8)
  }

  func preparePlayPause() {
    playPause = Material.Button(image: Icon.cm.pause, tintColor: Color.white)
    playPause.addTarget(self, action: #selector(self.tapPlayPause), for: .touchUpInside)

    layout(playPause).centerVertically().right(60).width(44).height(44)
  }

  func prepareSkip() {
    skip = Material.Button(image: Icon.cm.skipForward, tintColor: Color.white)
    skip.addTarget(self, action: #selector(self.tapSkip), for: .touchUpInside)

    layout(skip).centerVertically().right(16).width(44).height(44)
  }

  @objc
  func tapPlayPause(sender: UIView) {
    SpotifyPlayer.sharedInstance.togglePlayPause()
    self.playing = !self.playing
    self.refresh()
  }

  @objc
  func tapSkip(sender: UIView) {
    SpotifyPlayer.sharedInstance.skip()
  }

  func refresh() {
    playing = SpotifyPlayer.sharedInstance.isPlaying

    // hide the controls if there isn't a track in the queue
    if viewController.queue?.numTracks() == 0 {
      DispatchQueue.main.async {
        self.isHidden = true
      }
    } else {
      DispatchQueue.main.async {
        self.track.text = self.viewController.queue?.currentTrack().title()
        self.artist.text = self.viewController.queue?.currentTrack().artistDisplay()
        self.isHidden = false
      }
    }
  }
  
  func changeTrack() {
    self.refresh()
  }
  
  func changePlayPause() {
    self.refresh()
  }
}
