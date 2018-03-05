//
//  PlaylistsViewController.swift
//  QCast
//
//  Created by Vinnie Magro on 11/16/17.
//  Copyright © 2017 Vinnie Magro. All rights reserved.
//

import Foundation
import UIKit
import Material
import Kingfisher

import Libqcast

class PlaylistsViewController: TableViewController {
  private var menuButton: IconButton!

  private var playlistsService = MobileNewSpotifyPlaylistService()
  private var playlists: MobilePlaylistsList?

  private var queue: MobileQueueProtocol!

  open override func prepare() {
    super.prepare()
    prepareNavigationBar()
  }

  private func prepareNavigationBar() {
    navigationController!.navigationBar.tintColor = Color.white
    navigationItem.backButton.tintColor = Color.white
  }

  override func viewDidAppear(_ animated: Bool) {
    super.viewDidAppear(animated)

    // load the playlists
    playlistsService?.fetchPlaylists(self)
  }
}

extension PlaylistsViewController: UITextFieldDelegate {
  func textFieldDidEndEditing(_ textField: UITextField) {
    textField.resignFirstResponder()
  }

  func textFieldShouldReturn(_ textField: UITextField) -> Bool {
    textField.endEditing(true)
    return false
  }
}

extension PlaylistsViewController: MobilePlaylistsCallbackProtocol {

  func playlistsReceived(_ playlists: MobilePlaylistsList!) {
    self.playlists = playlists
    DispatchQueue.main.async {
      self.tableView.reloadData()
    }
  }

}

extension PlaylistsViewController {
  override func numberOfSections(in tableView: UITableView) -> Int {
    return 1
  }

  override func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
    if playlists != nil {
      return playlists!.numPlaylists()
    }
    return 0
  }

  override func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
    //        let cell = tableView.dequeueReusableCell(withIdentifier: "PlaylistsTrackTableViewCell", for: indexPath) as! PlaylistsTrackTableViewCell

    let cell = tableView.dequeueReusableCell(withIdentifier: "TableViewCell", for: indexPath) as! TableViewCell

    guard let playlists = playlists else {
      return cell
    }

    let playlist: LibqcastPlaylist
    do {
      try playlist = playlists.playlist(at: indexPath.row)
    } catch {
      return cell
    }

    cell.textLabel?.text = playlist.name()
    cell.imageView?.kf.setImage(with: URL(string: playlist.imageURL()), placeholder: Icon.cm.audioLibrary?.resize(toHeight: 64)?.resize(toWidth: 64))
    cell.detailTextLabel?.text = "\(playlist.numTracks()) tracks"
    cell.dividerColor = Color.grey.lighten2

    return cell
  }
  
  func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
    // open the selected playlist
    do {
      let playlist: LibqcastPlaylist
      try playlist = playlists!.playlist(at: indexPath.row)
      
      let vc = PlaylistViewController(playlist: playlist)
      navigationController?.pushViewController(vc, animated: true)
    } catch {
      // TODO: do something smart here
    }
  }
}
