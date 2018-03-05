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

@class FBFramebufferSurface;
@protocol FBControlCoreLogger;

/**
 A Video Encoder using SimDisplayVideoWriter.
 */
@interface FBVideoEncoderSimulatorKit : NSObject

#pragma mark Initializers

/**
 Create a new Encoder with the provided parameters.

 @param surface the surface to connect to.
 @param videoPath the video path to write to.
 @param logger the optional logger to log to.
 @return a new Encoder Instance.
 */
+ (instancetype)encoderWithRenderable:(FBFramebufferSurface *)surface videoPath:(NSString *)videoPath logger:(nullable id<FBControlCoreLogger>)logger;

#pragma mark Public Methods

/**
 YES if this class is supported, NO otherwise.
 */
+ (BOOL)isSupported;

/**
 Starts Recording Video.

 @return a future that resolves when the recording starts.
 */
- (FBFuture<NSNull *> *)startRecording;

/**
 Stops Recording Video.

 @return a future that resolves when the recording stops.
 */
- (FBFuture<NSNull *> *)stopRecording;

#pragma mark Properties

/**
 The Queue used for Serializing Media Actions.
 */
@property (nonatomic, strong, readonly) dispatch_queue_t mediaQueue;

@end

NS_ASSUME_NONNULL_END
