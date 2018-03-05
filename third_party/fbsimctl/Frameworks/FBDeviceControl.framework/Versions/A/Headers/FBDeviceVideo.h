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

@class AVCaptureSession;
@class FBDevice;

/**
 A Class for obtaining Video Configuration for a Device.
 */
@interface FBDeviceVideo : NSObject <FBVideoRecordingSession>

#pragma mark Initializers

/**
 Obtains the AVCaptureSession for a Device.

 @param device the Device to obtain the Session for.
 @param error an error out for any error that occurs.
 @return A Capture Session if successful, nil otherwise.
 */
+ (nullable AVCaptureSession *)captureSessionForDevice:(FBDevice *)device error:(NSError **)error;

/**
 A Factory method for obtaining the Video for a Device.

 @param device the Device.
 @param filePath the location of the video to record to, will be deleted if it already exists.
 @param error an error out for any error that occurs.
 */
+ (nullable instancetype)videoForDevice:(FBDevice *)device filePath:(NSString *)filePath error:(NSError **)error;

#pragma mark Public

/**
 Starts Recording the Video for a Device.

 @return a Future that resolves when recording has started.
 */
- (FBFuture<NSNull *> *)startRecording;

/**
 Stops Recording the Video for a Device.

 @return a Future that resolves when recording has stopped.
 */
- (FBFuture<NSNull *> *)stopRecording;

@end

NS_ASSUME_NONNULL_END
