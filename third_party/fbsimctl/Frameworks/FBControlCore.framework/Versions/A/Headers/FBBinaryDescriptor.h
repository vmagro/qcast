/**
 * Copyright (c) 2015-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#import <Foundation/Foundation.h>

#import <FBControlCore/FBJSONConversion.h>

NS_ASSUME_NONNULL_BEGIN

/**
 Concrete value wrapper around a binary artifact.
 */
@interface FBBinaryDescriptor : NSObject <NSCopying, FBJSONSerializable, FBJSONDeserializable>

/**
 The Designated Initializer.

 @param name The name of the executable. Must not be nil.
 @param path The path to the executable. Must not be nil.
 @param architectures The supported architectures of the executable. Must not be nil.
 @return a new FBBinaryDescriptor instance.
 */
- (instancetype)initWithName:(NSString *)name path:(NSString *)path architectures:(NSSet *)architectures;

/**
 An initializer for FBBinaryDescriptor that checks the nullability of the arguments

 @param name The name of the executable. May be nil.
 @param path The path to the executable. May be nil.
 @param architectures The supported architectures of the executable. Must not be nil.
 @return a new FBBinaryDescriptor instance, if all arguments are non-nil. Nil otherwise.
 */
+ (nullable instancetype)withName:(NSString *)name path:(NSString *)path architectures:(NSSet *)architectures;

/**
 The Name of the Executable.
 */
@property (nonatomic, copy, readonly) NSString *name;

/**
 The File Path to the Executable.
 */
@property (nonatomic, copy, readonly) NSString *path;

/**
 The Supported Architectures of the Executable.
 */
@property (nonatomic, copy, readonly) NSSet *architectures;

@end

/**
 Conveniences for building FBBinaryDescriptor instances
 */
@interface FBBinaryDescriptor (Helpers)

/**
 Returns the FBBinaryDescriptor for the given binary path
 */
+ (nullable instancetype)binaryWithPath:(NSString *)path error:(NSError **)error;

@end

NS_ASSUME_NONNULL_END
