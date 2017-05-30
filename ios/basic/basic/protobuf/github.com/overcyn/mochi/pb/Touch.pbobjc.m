// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: github.com/overcyn/mochi/pb/touch.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers_RuntimeSupport.h>
#else
 #import "GPBProtocolBuffers_RuntimeSupport.h"
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/Any.pbobjc.h>
 #import <Protobuf/Timestamp.pbobjc.h>
#else
 #import "google/protobuf/Any.pbobjc.h"
 #import "google/protobuf/Timestamp.pbobjc.h"
#endif

 #import "github.com/overcyn/mochi/pb/Touch.pbobjc.h"
 #import "github.com/overcyn/mochi/pb/Layout.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - MochiPBTouchRoot

@implementation MochiPBTouchRoot

// No extensions in the file and none of the imports (direct or indirect)
// defined extensions, so no need to generate +extensionRegistry.

@end

#pragma mark - MochiPBTouchRoot_FileDescriptor

static GPBFileDescriptor *MochiPBTouchRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"mochi.touch"
                                                 objcPrefix:@"MochiPB"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - MochiPBRecognizer

@implementation MochiPBRecognizer

@dynamic id_p;
@dynamic name;
@dynamic hasRecognizer, recognizer;

typedef struct MochiPBRecognizer__storage_ {
  uint32_t _has_storage_[1];
  NSString *name;
  GPBAny *recognizer;
  int64_t id_p;
} MochiPBRecognizer__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "id_p",
        .dataTypeSpecific.className = NULL,
        .number = MochiPBRecognizer_FieldNumber_Id_p,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MochiPBRecognizer__storage_, id_p),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "name",
        .dataTypeSpecific.className = NULL,
        .number = MochiPBRecognizer_FieldNumber_Name,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MochiPBRecognizer__storage_, name),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "recognizer",
        .dataTypeSpecific.className = GPBStringifySymbol(GPBAny),
        .number = MochiPBRecognizer_FieldNumber_Recognizer,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(MochiPBRecognizer__storage_, recognizer),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MochiPBRecognizer class]
                                     rootClass:[MochiPBTouchRoot class]
                                          file:MochiPBTouchRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MochiPBRecognizer__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MochiPBRecognizerList

@implementation MochiPBRecognizerList

@dynamic recognizersArray, recognizersArray_Count;

typedef struct MochiPBRecognizerList__storage_ {
  uint32_t _has_storage_[1];
  NSMutableArray *recognizersArray;
} MochiPBRecognizerList__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "recognizersArray",
        .dataTypeSpecific.className = GPBStringifySymbol(MochiPBRecognizer),
        .number = MochiPBRecognizerList_FieldNumber_RecognizersArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MochiPBRecognizerList__storage_, recognizersArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MochiPBRecognizerList class]
                                     rootClass:[MochiPBTouchRoot class]
                                          file:MochiPBTouchRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MochiPBRecognizerList__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MochiPBTapRecognizer

@implementation MochiPBTapRecognizer

@dynamic count;
@dynamic recognizedFunc;

typedef struct MochiPBTapRecognizer__storage_ {
  uint32_t _has_storage_[1];
  int64_t count;
  int64_t recognizedFunc;
} MochiPBTapRecognizer__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "count",
        .dataTypeSpecific.className = NULL,
        .number = MochiPBTapRecognizer_FieldNumber_Count,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MochiPBTapRecognizer__storage_, count),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "recognizedFunc",
        .dataTypeSpecific.className = NULL,
        .number = MochiPBTapRecognizer_FieldNumber_RecognizedFunc,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MochiPBTapRecognizer__storage_, recognizedFunc),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MochiPBTapRecognizer class]
                                     rootClass:[MochiPBTouchRoot class]
                                          file:MochiPBTouchRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MochiPBTapRecognizer__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\001\002\016\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MochiPBTapEvent

@implementation MochiPBTapEvent

@dynamic hasTimestamp, timestamp;
@dynamic hasPosition, position;

typedef struct MochiPBTapEvent__storage_ {
  uint32_t _has_storage_[1];
  GPBTimestamp *timestamp;
  MochiPBPoint *position;
} MochiPBTapEvent__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "timestamp",
        .dataTypeSpecific.className = GPBStringifySymbol(GPBTimestamp),
        .number = MochiPBTapEvent_FieldNumber_Timestamp,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MochiPBTapEvent__storage_, timestamp),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "position",
        .dataTypeSpecific.className = GPBStringifySymbol(MochiPBPoint),
        .number = MochiPBTapEvent_FieldNumber_Position,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MochiPBTapEvent__storage_, position),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MochiPBTapEvent class]
                                     rootClass:[MochiPBTouchRoot class]
                                          file:MochiPBTouchRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MochiPBTapEvent__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MochiPBPressRecognizer

@implementation MochiPBPressRecognizer


typedef struct MochiPBPressRecognizer__storage_ {
  uint32_t _has_storage_[1];
} MochiPBPressRecognizer__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MochiPBPressRecognizer class]
                                     rootClass:[MochiPBTouchRoot class]
                                          file:MochiPBTouchRoot_FileDescriptor()
                                        fields:NULL
                                    fieldCount:0
                                   storageSize:sizeof(MochiPBPressRecognizer__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MochiPBPanRecognizer

@implementation MochiPBPanRecognizer


typedef struct MochiPBPanRecognizer__storage_ {
  uint32_t _has_storage_[1];
} MochiPBPanRecognizer__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MochiPBPanRecognizer class]
                                     rootClass:[MochiPBTouchRoot class]
                                          file:MochiPBTouchRoot_FileDescriptor()
                                        fields:NULL
                                    fieldCount:0
                                   storageSize:sizeof(MochiPBPanRecognizer__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
