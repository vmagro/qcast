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

NS_ASSUME_NONNULL_BEGIN

/**
 A Class for handling Framework Loading of Private Frameworks that FBSimulatorControl depends on.
 */
@interface FBSimulatorControlFrameworkLoader : FBControlCoreFrameworkLoader

/**
 The Frameworks needed for most operations.
 */
@property (nonatomic, strong, class, readonly) FBSimulatorControlFrameworkLoader *essentialFrameworks;

/**
 All of the Frameworks for operations involving the HID and Framebuffer.
 */
@property (nonatomic, strong, class, readonly) FBSimulatorControlFrameworkLoader *xcodeFrameworks;

@end

NS_ASSUME_NONNULL_END
