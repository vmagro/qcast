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

/**
 Encodes Device Video to a File, using an AVCaptureSession
 */
@interface FBDeviceVideoFileEncoder : NSObject

#pragma mark Initializers

/**
 Creates a Video Encoder with the provided Parameters.

 @param session the Session to record from.
 @param filePath the File Path to record to.
 @param logger the logger to use.
 @param error an error out for any error that occurs.
 */
+ (nullable instancetype)encoderWithSession:(AVCaptureSession *)session filePath:(NSString *)filePath logger:(id<FBControlCoreLogger>)logger error:(NSError **)error;

#pragma mark Public Methods

/**
 Starts the Video Encoder.

 @return A future that resolves when encoding has started.
 */
- (FBFuture<NSNull *> *)startRecording;

/**
 Stops the Video Encoder.
 If the encoder is running, it will block until the Capture Session has been torn down.

 @return A future that resolves when encoding has stopped.
 */
- (FBFuture<NSNull *> *)stopRecording;

@end

NS_ASSUME_NONNULL_END
