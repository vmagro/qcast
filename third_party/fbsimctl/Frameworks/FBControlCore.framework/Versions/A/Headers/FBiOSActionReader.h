/**
 * Copyright (c) 2015-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#import <Foundation/Foundation.h>

#import <FBControlCore/FBEventReporter.h>
#import <FBControlCore/FBiOSTargetAction.h>
#import <FBControlCore/FBTerminationHandle.h>

NS_ASSUME_NONNULL_BEGIN

@class FBUploadHeader;
@class FBUploadedDestination;
@class FBiOSActionRouter;

@protocol FBiOSTarget;
@protocol FBiOSTargetAction;
@protocol FBiOSActionReaderDelegate;

/**
 The Termination Handle Type for an Action Reader.
 */
extern FBTerminationHandleType const FBTerminationHandleTypeActionReader;

/**
 Routes an Actions for Sockets and Files.
 */
@interface FBiOSActionReader : NSObject <FBTerminationAwaitable>

/**
 Initializes an Action Reader for a target, on a socket.
 The default routing of the target will be used.

 @param target the target to run against.
 @param delegate the delegate to notify.
 @param port the port to bind on.
 @return a Socket Reader.
 */
+ (instancetype)socketReaderForTarget:(id<FBiOSTarget>)target delegate:(id<FBiOSActionReaderDelegate>)delegate port:(in_port_t)port;


/**
 Initializes an Action Reader for a router, on a socket.
 The Designated Initializer.

 @param router the router to use.
 @param delegate the delegate to notify.
 @param port the port to bind on.
 @return a Socket Reader.
 */
+ (instancetype)socketReaderForRouter:(FBiOSActionRouter *)router delegate:(id<FBiOSActionReaderDelegate>)delegate port:(in_port_t)port;

/**
 Initializes an Action Reader for a router, between file handles.
 The default routing of the target will be used.

 @param target the target to run against.
 @param delegate the delegate to notify.
 @param readHandle the handle to read.
 @param writeHandle the handle to write to.
 @return a Socket Reader.
 */
+ (instancetype)fileReaderForTarget:(id<FBiOSTarget>)target delegate:(id<FBiOSActionReaderDelegate>)delegate readHandle:(NSFileHandle *)readHandle writeHandle:(NSFileHandle *)writeHandle;

/**
 Initializes an Action Reader for a router, between file handles.

 @param router the router to use.
 @param delegate the delegate to notify.
 @param readHandle the handle to read.
 @param writeHandle the handle to write to.
 @return a Socket Reader.
 */
+ (instancetype)fileReaderForRouter:(FBiOSActionRouter *)router delegate:(id<FBiOSActionReaderDelegate>)delegate readHandle:(NSFileHandle *)readHandle writeHandle:(NSFileHandle *)writeHandle;

/**
 Create and Listen to the socket.

 @param error an error out for any error that occurs.
 @return YES if successful, NO otherwise.
 */
- (BOOL)startListeningWithError:(NSError **)error;

/**
 Stop listening to the socket

 @param error an error out for any error that occurs.
 @return YES if successful, NO otherwise.
 */
- (BOOL)stopListeningWithError:(NSError **)error;

@end

/**
 The Delegate for the Action Reader.
 */
@protocol FBiOSActionReaderDelegate <FBiOSTargetActionDelegate, FBEventReporter>

/**
 Called when the Reader has finished reading.

 @param reader the reader.
 */
- (void)readerDidFinishReading:(FBiOSActionReader *)reader;

/**
 Called when the Reader failed to interpret some input.

 @param reader the reader.
 @param input the line of input
 @param error the generated error.
 */
- (nullable NSString *)reader:(FBiOSActionReader *)reader failedToInterpretInput:(NSString *)input error:(NSError *)error;

/**
 Called when the Reader failed to interpret some input.

 @param reader the reader.
 @param header the header of the file being uploaded.
 @return the string to write back to the reader, if relevant.
 */
- (nullable NSString *)reader:(FBiOSActionReader *)reader willStartReadingUpload:(FBUploadHeader *)header;

/**
 Called when the Reader failed to interpret some input.

 @param reader the reader.
 @param destination the destination of the upload.
 @return the string to write back to the reader, if relevant.
 */
- (nullable NSString *)reader:(FBiOSActionReader *)reader didFinishUpload:(FBUploadedDestination *)destination;

/**
 Called when the Reader is about to perform an action.

 @param reader the reader performing the action
 @param action the action to be performed
 @param target the target
 @return the string to write back to the reader, if relevant.
 */
- (nullable NSString *)reader:(FBiOSActionReader *)reader willStartPerformingAction:(id<FBiOSTargetAction>)action onTarget:(id<FBiOSTarget>)target;

/**
 Called when the Reader has successfully performed an action

 @param reader the reader performing the action
 @param action the action to be performed
 @param target the target
 @return the string to write back to the reader, if relevant.
*/
- (nullable NSString *)reader:(FBiOSActionReader *)reader didProcessAction:(id<FBiOSTargetAction>)action onTarget:(id<FBiOSTarget>)target;

/**
 Called when the Reader has failed to perform an action

 @param reader the reader performing the action
 @param action the action to be performed
 @param target the target
 @param error the error.
 @return the string to write back to the reader, if relevant.
 */
- (nullable NSString *)reader:(FBiOSActionReader *)reader didFailToProcessAction:(id<FBiOSTargetAction>)action onTarget:(id<FBiOSTarget>)target error:(NSError *)error;

@end

NS_ASSUME_NONNULL_END
