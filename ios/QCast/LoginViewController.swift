//
//  LoginViewController.swift
//  QCast
//
//  Created by Vinnie Magro on 11/10/17.
//  Copyright © 2017 Vinnie Magro. All rights reserved.
//

import UIKit
import Material
import SafariServices
import Libqcast

protocol LoginViewControllerDelegate {
  func loginDidSucceed()
  func loginCanceled()
}

class LoginViewController: UIViewController {
  fileprivate var menuButton: IconButton!

  fileprivate var loginButton: Button!
  fileprivate var safari: SFSafariViewController?
  
  fileprivate var navContainer = UIView()
  fileprivate var descriptionLabel: UILabel!
  
  var delegate: LoginViewControllerDelegate?
  
  var instructions: String? {
    didSet {
      if descriptionLabel != nil {
        descriptionLabel.text = instructions
      }
    }
  }

  open override func viewDidLoad() {
    super.viewDidLoad()
    view.backgroundColor = Color.grey.lighten5

    prepareNavContainer()
    prepareLoginButton()
    
    let service = LibqcastAuthService()
    if service?.loggedIn() ?? false {
      self.safari?.dismiss(animated: true, completion: nil)
      delegate?.loginDidSucceed()
    } else {
      LibqcastAuthService().wait(forLogin: self)
    }
  }
  
  override func viewDidAppear(_ animated: Bool) {
    super.viewDidAppear(animated)
    
    let service = LibqcastAuthService()
    if service?.loggedIn() ?? false {
      self.safari?.dismiss(animated: true, completion: nil)
      delegate?.loginDidSucceed()
    } else {
      LibqcastAuthService().wait(forLogin: self)
    }
  }
  
  override var preferredStatusBarStyle: UIStatusBarStyle {
    return .lightContent
  }
}

fileprivate extension LoginViewController {
  
  func prepareNavContainer() {
    navContainer.backgroundColor = Color.green.base
    
    view.layout(navContainer).top(0).left(0).right(0).height(192)
    
    prepareTitle()
    prepareDescription()
    prepareCancelButton()
  }
  
  func prepareCancelButton() {
    let cancelButton = Button(title: "Cancel", titleColor: Color.white)
    cancelButton.addTarget(self, action: #selector(self.cancelPressed), for: .touchUpInside)
    view.layout(cancelButton).top(20).left(14).height(44)
  }
  
  func prepareTitle() {
    let titleLabel = UILabel(frame: CGRect(x: 0, y: 0, width: 200, height: 44))
    titleLabel.text = "Login"
    titleLabel.textColor = Color.white
    titleLabel.font = UIFont.systemFont(ofSize: 17, weight: .semibold)
    
    navContainer.layout(titleLabel).centerHorizontally().top(32)
  }
  
  func prepareDescription() {
    descriptionLabel = UILabel()
    descriptionLabel.text = instructions
    descriptionLabel.textColor = Color.white
    descriptionLabel.font = RobotoFont.light(with: 22)
    descriptionLabel.backgroundColor = Color.green.base
    descriptionLabel.numberOfLines = 0
    
    navContainer.layout(descriptionLabel).left(8).right(8).bottom(16)
  }

  func prepareLoginButton() {
    loginButton = RaisedButton(title: "Sign in with Spotify", titleColor: Color.white)
    
    loginButton.backgroundColor = Color.green.base
    loginButton.cornerRadiusPreset = .cornerRadius3
    loginButton.depthPreset = .depth3
    loginButton.titleLabel?.font = RobotoFont.light(with: 22)
    
    loginButton.addTarget(self, action: #selector(self.handleLoginButton), for: .touchUpInside)
    
    self.view.layout(loginButton).center().width(200).height(44)
  }

}

extension LoginViewController {
  @objc
  func handleLoginButton(sender: Button) {
    let delegate = UIApplication.shared.delegate as! AppDelegate
    
    // give preference to opening the Spotify app, but if we can't then open in a Safari view
    if SPTAuth.supportsApplicationAuthentication() {
      let loginUrl = delegate.sptAuth.spotifyAppAuthenticationURL()!
      UIApplication.shared.open(loginUrl, options: [:], completionHandler: nil)
    } else {
      let loginUrl = delegate.sptAuth.spotifyWebAuthenticationURL()!
      safari = SFSafariViewController(url: loginUrl)
      self.present(safari!, animated: true, completion: {
        //TODO: do something here?
      })
    }
  }
  
  @objc
  func cancelPressed(sender: Button) {
    delegate?.loginCanceled()
  }
}

extension LoginViewController: LibqcastAuthCallbackProtocol {
  func loginSucceeded() {
    safari?.dismiss(animated: true, completion: {
      // tell the presenting view controller that we finished
      self.delegate?.loginDidSucceed()
    })
  }
  
  func loginFailed(_ err: Error!) {
    // TODO: handle this
  }
}
