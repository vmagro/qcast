//
//  PartyMoodViewController.swift
//  QCastDev
//
//  Created by Vinnie Magro on 11/22/17.
//

import Foundation
import UIKit
import Material
import Libqcast

class PartyMoodViewController: UIViewController {
  
  fileprivate var instructionsContainer: UIView!
  fileprivate var nextButton: Button!
  
  open override func viewDidLoad() {
    super.viewDidLoad()
    view.backgroundColor = Color.grey.lighten5
    
    prepareNavigationItem()
    prepareInstructions()
    prepareNextButton()
    
    // kick off a load to get Spotify playlists
    MobileNewSpotifyPlaylistService().fetchPlaylists(self)
  }
  
  open override func viewWillAppear(_ animated: Bool) {
    super.viewWillAppear(animated)
    
    prepareNavigationItem()
  }
}

fileprivate extension PartyMoodViewController {
  
  func prepareNavigationItem() {
    // reset navigation bar from the clear crap in the home screen
    navigationItem.titleLabel.text = "Set The Mood"
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
    label.text = "Choose a playlist to kick off the mood of your party. (optional)"
    label.textColor = Color.white
    label.font = RobotoFont.light(with: 22)
    label.numberOfLines = 0
    
    view.layout(instructionsContainer).left(0).top(0).right(0).height(128)
    instructionsContainer.layout(label).left(8).right(8).bottom(16)
  }
  
  func prepareNextButton() {
    nextButton = Button()
    
    nextButton.title = "Continue"
    nextButton.addTarget(self, action: #selector(self.handleNextButton), for: .touchUpInside)
    
    view.layout(nextButton).bottom(14).right(12)
  }
  
}

fileprivate extension PartyMoodViewController {
  @objc
  func handleNextButton() {
    navigationController?.pushViewController(PartyConfirmViewController(), animated: true)
  }
  
}

extension PartyMoodViewController: MobilePlaylistsCallbackProtocol {
  func playlistsReceived(_ playlists: MobilePlaylistsList!) {
    NSLog("party mood got %d playlists", playlists.numPlaylists())
  }
}
