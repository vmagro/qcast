/**
 * Copyright (c) 2015-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#import <Foundation/Foundation.h>

#import <FBControlCore/FBControlCore.h>

/**
 Framework and Class Loading for XCTestBoostrap.
 */
@interface XCTestBootstrapFrameworkLoader : FBControlCoreFrameworkLoader

/**
 All of the Frameworks for XCTestBootstrap.
 */
@property (nonatomic, strong, class, readonly) XCTestBootstrapFrameworkLoader *allDependentFrameworks;

@end
