//
//  PartyConfirmViewController.swift
//  QCastDev
//
//  Created by Vinnie Magro on 11/24/17.
//

import Foundation
import UIKit
import Material
import CoreLocation
import Libqcast

class PartyConfirmViewController: UIViewController {
  
  // in progress party being created
  var partyName: String?
  var partyLocation: CLLocation?
  
  fileprivate var instructionsContainer: UIView!
  fileprivate var nextButton: Button!
  
  open override func viewDidLoad() {
    super.viewDidLoad()
    view.backgroundColor = Color.grey.lighten5
    
    prepareNavigationItem()
    prepareInstructions()
    prepareNextButton()
  }
  
  open override func viewWillAppear(_ animated: Bool) {
    super.viewWillAppear(animated)
    
    prepareNavigationItem()
  }
}

fileprivate extension PartyConfirmViewController {
  
  func prepareNavigationItem() {
    // reset navigation bar from the clear crap in the home screen
    navigationItem.titleLabel.text = "Start the Party"
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
    label.text = "Does all this look correct?"
    label.textColor = Color.white
    label.font = RobotoFont.light(with: 22)
    label.numberOfLines = 0
    
    view.layout(instructionsContainer).left(0).top(0).right(0).height(128)
    instructionsContainer.layout(label).left(8).right(8).bottom(16)
  }
  
  func prepareNextButton() {
    nextButton = Button()
    
    nextButton.title = "Start!"
    nextButton.addTarget(self, action: #selector(self.handleNextButton), for: .touchUpInside)
    
    view.layout(nextButton).bottom(14).right(12)
  }
  
}

fileprivate extension PartyConfirmViewController {
  @objc
  func handleNextButton() {
    let coord = partyLocation?.coordinate
    var loc: LibqcastLocation? = nil
    if coord != nil {
      loc = LibqcastNewLocation(coord!.latitude, coord!.longitude)
    }
    do {
      let service = LibqcastRethinkPartyService()!
      try service.startParty(self.partyName, location: loc, moodPlaylist: nil)
    } catch {
      let alert = UIAlertController(title: "Oops", message: "Something happened and we weren't able to create your party.\nPlease try again later.", preferredStyle: .alert)
      alert.addAction(UIAlertAction(title: "Close", style: .default, handler: { (_ action) in
        alert.dismiss(animated: true, completion: nil)
      }))
      self.present(alert, animated: true, completion: nil)
    }
  }
  
}
