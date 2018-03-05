/**
 * Copyright (c) 2015-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#import <Foundation/Foundation.h>

#import <FBControlCore/FBFuture.h>
#import <FBControlCore/FBiOSTargetCommandForwarder.h>
#import <FBControlCore/FBTerminationHandle.h>

NS_ASSUME_NONNULL_BEGIN

@protocol FBFileConsumer;
@protocol FBVideoRecordingSession;

/**
 The Termination Handle Type for an Recording Operation.
 */
extern FBTerminationHandleType const FBTerminationTypeHandleVideoRecording;

/**
 Defines an async Video Recording Operation, suitable for being stopped.
 */
@protocol FBVideoRecordingSession <NSObject, FBTerminationHandle>

@end

/**
 Defines an interface for Video Recording.
 */
@protocol FBVideoRecordingCommands <NSObject, FBiOSTargetCommand>

/**
 Starts the Recording of Video to a File.

 @param filePath an optional filePath to write to. If not provided, a default file path will be used.
 @return A Future, wrapping the recording session.
 */
- (FBFuture<id<FBVideoRecordingSession>> *)startRecordingToFile:(nullable NSString *)filePath;

/**
 Stops the Recording of Video.

 @return A Future, resolved when recording has stopped
 */
- (FBFuture<NSNull *> *)stopRecording;

@end

NS_ASSUME_NONNULL_END
