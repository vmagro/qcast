//
//  NamePartyViewController.swift
//  QCastDev
//
//  Created by Vinnie Magro on 11/21/17.
//

import Foundation
import UIKit
import Material
import Libqcast

class NamePartyViewController: UIViewController {

  fileprivate var instructionsContainer: UIView!
  fileprivate var nameInput: TextField!
  
  // allow us to force not presenting the login view controller so that we can pop back safely
  fileprivate var forceNoLogin = false
  fileprivate var loginController: LoginViewController?

  open override func viewDidLoad() {
    super.viewDidLoad()
    view.backgroundColor = Color.grey.lighten5

    prepareNavigationItem()
    prepareInstructions()
    prepareTextField()
  }

  open override func viewWillAppear(_ animated: Bool) {
    super.viewWillAppear(animated)

    prepareNavigationItem()
  }
  
  override func viewDidAppear(_ animated: Bool) {
    super.viewDidAppear(animated)
    
    if !LibqcastAuthService().loggedIn() && !forceNoLogin {
      loginController = LoginViewController()
      loginController?.delegate = self
      loginController?.instructions = "Login with Spotify to start a party and play music!"
      
      present(loginController!, animated: true, completion: {})
    }
    
    _ = self.nameInput.becomeFirstResponder()
  }
}

extension NamePartyViewController: LoginViewControllerDelegate {
  func loginDidSucceed() {
    loginController?.dismiss(animated: true, completion: {})
  }
  
  func loginCanceled() {
    forceNoLogin = true
    loginController?.dismiss(animated: true, completion: {
      self.navigationController?.popViewController(animated: true)
    })
  }
}

fileprivate extension NamePartyViewController {

  func prepareNavigationItem() {
    // reset navigation bar from the clear crap in the home screen
    navigationItem.titleLabel.text = "Name Your Party"
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
    label.text = "Give your party a name so that your guests can find it more easily!"
    label.textColor = Color.white
    label.font = RobotoFont.light(with: 22)
    label.numberOfLines = 0

    view.layout(instructionsContainer).left(0).top(0).right(0).height(128)
    instructionsContainer.layout(label).left(8).right(8).bottom(16)
  }

  func prepareTextField() {
    nameInput = TextField()
    nameInput.delegate = self
    
    nameInput.backgroundColor = Color.white
    nameInput.placeholder = "Party Name"
    nameInput.placeholderVerticalOffset = 12
    nameInput.cornerRadiusPreset = .cornerRadius1
    nameInput.textInset = 8
    
    nameInput.returnKeyType = .next
    
    view.layout(nameInput).top(160).left(8).right(8).height(48)
  }
}

extension NamePartyViewController : UITextFieldDelegate {
  
  func textFieldShouldReturn(_ textField: UITextField) -> Bool {
    // only allow going to the next page if there is text in the box
    if textField.text!.count > 0 {
      textField.resignFirstResponder()
      self.handleNextButton()
      return true
    }
    return false
  }
  
}

fileprivate extension NamePartyViewController {
  @objc
  func handleNextButton() {
    let vc = LocationViewController()
    vc.partyName = self.nameInput.text
    navigationController?.pushViewController(vc, animated: true)
  }

}
