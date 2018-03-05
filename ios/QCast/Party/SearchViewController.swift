//
//  SearchViewController.swift
//  QCast
//
//  Created by Vinnie Magro on 11/9/17.
//  Copyright © 2017 Vinnie Magro. All rights reserved.
//

import Foundation
import UIKit
import Material
import Kingfisher

import Libqcast

class SearchViewController: UIViewController {
  private var searchBar: UISearchBar!

  private weak var searchDebounce: Timer?
  private var searcher = MobileSpotifySearchService()

  private var queue: MobileQueueProtocol!
  private var tableView: TrackListTable!
  
  override func viewDidLoad() {
    super.viewDidLoad()
    prepare()
  }
  
  open override func viewDidAppear(_ animated: Bool) {
    super.viewDidAppear(animated)

    searchBar.becomeFirstResponder()
  }

  open func prepare() {
    prepareResults()
    prepareNavigationBar()
    prepareSearchBar()
  }
  
  private func prepareResults() {
    tableView = TrackListTable()
    view.layout(tableView).left(0).right(0).top(0).bottom(0)
  }

  private func prepareNavigationBar() {
    navigationController!.navigationBar.tintColor = Color.white
    navigationItem.backButton.tintColor = Color.white
  }

  private func prepareSearchBar() {
    // try to fill up the whole screen width but leave some space for the back button and also for right padding
    let width = UIScreen.main.bounds.width - 48 - 32
    searchBar = UISearchBar(frame: CGRect(x: 0, y: 0, width: width, height: 64))
    searchBar.sizeToFit()
    searchBar.placeholder = "Search"
    searchBar.delegate = self
    searchBar.returnKeyType = .search
    // make the cursor come back as blue
    searchBar.tintColor = Color.blue
    navigationItem.rightBarButtonItem = UIBarButtonItem(customView: searchBar)
  }

  @objc
  func search() {
    searcher?.search(searchBar.text, callback: self)
  }
}

extension SearchViewController: UISearchBarDelegate {
  func searchBar(_ searchBar: UISearchBar, textDidChange searchText: String) {
    NSObject.cancelPreviousPerformRequests(withTarget: self, selector: #selector(self.search), object: nil)
    self.perform(#selector(self.search), with: nil, afterDelay: 0.2)
  }
  
  func searchBarCancelButtonClicked(_ searchBar: UISearchBar) {
    NSObject.cancelPreviousPerformRequests(withTarget: self, selector: #selector(self.search), object: nil)
    self.perform(#selector(self.search), with: nil, afterDelay: 0.2)
  }
  
  func searchBarSearchButtonClicked(_ searchBar: UISearchBar) {
    searchBar.resignFirstResponder()
  }
  
  func searchBarTextDidEndEditing(_ searchBar: UISearchBar) {
    searchBar.resignFirstResponder()
  }
}

extension SearchViewController: MobileSearchCallbackProtocol {

  func searchResultsReceived(_ results: MobileSearchResults!) {
    print("Loaded \(results.numTracks()) tracks")
    
    // convert the results to a list of tracks
    var tracks = [LibqcastTrack]()
    for i in 0..<results.numTracks() {
      do {
        let track: LibqcastTrack
        try track = results.track(at: i)
        tracks.append(track)
      } catch {
        // TODO: do something smart
      }
    }
    self.tableView.tracks = tracks
  }

}
