//
//  AppNavigationController.swift
//  QCast
//
//  Created by Vinnie Magro on 11/9/17.
//  Copyright © 2017 Vinnie Magro. All rights reserved.
//

import Foundation
import UIKit
import Material

class AppNavigationController: NavigationController {
  open override func prepare() {
    super.prepare()

    guard let v = navigationBar as? NavigationBar else {
      return
    }

    v.backgroundColor = Color.green.base
    v.tintColor = Color.white

    v.depthPreset = .none
    v.dividerColor = Color.grey.lighten2
    
    // fix the back button being way too far to the right
    v.contentEdgeInsets = UIEdgeInsets(top: 0, left: -8, bottom: 0, right: 0)
  }
}
