//
//  PlaylistViewController.swift
//  QCastDev
//
//  Created by Vinnie Magro on 11/27/17.
//

import Foundation
import UIKit
import Libqcast

class PlaylistViewController: UIViewController {
  var tracksTable: TrackListTable!
  var playlist: LibqcastPlaylist!
  
  init(playlist: LibqcastPlaylist) {
    super.init(nibName: nil, bundle: nil)
    
    self.playlist = playlist
  }
  
  required init?(coder aDecoder: NSCoder) {
    super.init(coder: aDecoder)
  }
  
  override func viewDidLoad() {
    super.viewDidLoad()
    prepareTracksTable()
    
    // fire off a request to load the tracks in the playlist
    MobilePlaylistService().fetchPlaylistTracks(self.playlist, callback: self)
  }
  
  func prepareTracksTable() {
    tracksTable = TrackListTable()
    
    view.layout(tracksTable).left(0).right(0).top(0).bottom(0)
  }
  
}

extension PlaylistViewController: MobilePlaylistCallbackProtocol {
  func playlistTracksReceived(_ playlist: LibqcastPlaylist!) {
    // convert the tracks to an array and set them on the table view
  }
}
