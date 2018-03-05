//
//  AppDelegate.swift
//  QCast
//
//  Created by Vinnie Magro on 11/9/17.
//  Copyright © 2017 Vinnie Magro. All rights reserved.
//

import UIKit
import Material
import Libqcast

@UIApplicationMain
class AppDelegate: UIResponder, UIApplicationDelegate {

  var window: UIWindow?
  var sptAuth: SPTAuth!

  func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplicationLaunchOptionsKey: Any]?) -> Bool {
    window = UIWindow(frame: UIScreen.main.bounds)

    // setup native library and authentication framework before instantiating a view controller that may (read: does) use native functionality
    // set up qcast native library
    LibqcastInit()
    
    sptAuth = SPTAuth()
    sptAuth.clientID = LibqcastSpotifyClientID
    sptAuth.redirectURL = URL(string: LibqcastSpotifyRedirectURL)
    // use our exchange service
    sptAuth.tokenSwapURL = URL(string: LibqcastSpotifySwapURL)
    sptAuth.tokenRefreshURL = URL(string: LibqcastSpotifyRefreshURL)
    // get the required scopes from the native library
    sptAuth.requestedScopes = LibqcastSpotifyScopes().components(separatedBy: ",")
    // this allows SPTAuth to automatically save its session for future use
    sptAuth.sessionUserDefaultsKey = "SpotifySession"
    
    // if the session is not nil just renew it now so that we know it'll last longer
    if sptAuth.session != nil {
      // immediately login with the existing session and then fire off a request to renew the session
      handleAuthCallback(error: nil, session: sptAuth.session)
      sptAuth.renewSession(sptAuth.session, callback: self.handleAuthCallback)
    }
    // if we aren't already logged in we'll show the user the login view controller before they can do anything useful
    
    let homeViewController = AppNavigationController(rootViewController: HomeViewController())

    window!.rootViewController = homeViewController
    window!.makeKeyAndVisible()

    return true
  }

  func application(_ app: UIApplication,
    open url: URL,
  options: [UIApplicationOpenURLOptionsKey: Any] = [:]) -> Bool {
    if sptAuth.canHandle(url) {
      sptAuth.handleAuthCallback(withTriggeredAuthURL: url, callback: self.handleAuthCallback)
      return true
    }
    return false
  }
  
  func handleAuthCallback(error: Error?, session: SPTSession?) {
    print("Logged in with error = \(error), session = \(session)\n")
    if error != nil {
      // TODO: handle error better
      return
    }
    // login asynchronously since it requires some round trips
    DispatchQueue.global(qos: .background).async {
      do {
        try LibqcastAuthService().login(withToken: session?.accessToken)
        
        // setup the SpotifyPlayer so that it'll be ready by the time we need it
        if !SpotifyPlayer.sharedInstance.isLoggedIn {
          do {
            try SpotifyPlayer.sharedInstance.login(accessToken: session!.accessToken)
          } catch {
            // TODO: do something
          }
        }
      } catch {
        // TODO: handle error better
        return
      }
    }
  }
}
