// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gomatcha.io/matcha/pb/view/stacknav/stacknavigator.proto

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

 #import "gomatcha.io/matcha/pb/view/stacknav/Stacknavigator.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - MatchaStackScreenPBStacknavigatorRoot

@implementation MatchaStackScreenPBStacknavigatorRoot

// No extensions in the file and no imports, so no need to generate
// +extensionRegistry.

@end

#pragma mark - MatchaStackScreenPBStacknavigatorRoot_FileDescriptor

static GPBFileDescriptor *MatchaStackScreenPBStacknavigatorRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"matcha.view.stacknav"
                                                 objcPrefix:@"MatchaStackScreenPB"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - MatchaStackScreenPBChildView

@implementation MatchaStackScreenPBChildView

@dynamic viewId;
@dynamic barId;
@dynamic screenId;

typedef struct MatchaStackScreenPBChildView__storage_ {
  uint32_t _has_storage_[1];
  int64_t viewId;
  int64_t barId;
  int64_t screenId;
} MatchaStackScreenPBChildView__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "viewId",
        .dataTypeSpecific.className = NULL,
        .number = MatchaStackScreenPBChildView_FieldNumber_ViewId,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaStackScreenPBChildView__storage_, viewId),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "barId",
        .dataTypeSpecific.className = NULL,
        .number = MatchaStackScreenPBChildView_FieldNumber_BarId,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaStackScreenPBChildView__storage_, barId),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "screenId",
        .dataTypeSpecific.className = NULL,
        .number = MatchaStackScreenPBChildView_FieldNumber_ScreenId,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(MatchaStackScreenPBChildView__storage_, screenId),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaStackScreenPBChildView class]
                                     rootClass:[MatchaStackScreenPBStacknavigatorRoot class]
                                          file:MatchaStackScreenPBStacknavigatorRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaStackScreenPBChildView__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\003\001\006\000\002\005\000\003\010\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaStackScreenPBView

@implementation MatchaStackScreenPBView

@dynamic childrenArray, childrenArray_Count;

typedef struct MatchaStackScreenPBView__storage_ {
  uint32_t _has_storage_[1];
  NSMutableArray *childrenArray;
} MatchaStackScreenPBView__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "childrenArray",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaStackScreenPBChildView),
        .number = MatchaStackScreenPBView_FieldNumber_ChildrenArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaStackScreenPBView__storage_, childrenArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaStackScreenPBView class]
                                     rootClass:[MatchaStackScreenPBStacknavigatorRoot class]
                                          file:MatchaStackScreenPBStacknavigatorRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaStackScreenPBView__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaStackScreenPBBar

@implementation MatchaStackScreenPBBar

@dynamic title;
@dynamic backButtonHidden;
@dynamic customBackButtonTitle;
@dynamic backButtonTitle;
@dynamic titleViewId;
@dynamic rightViewIdsArray, rightViewIdsArray_Count;
@dynamic leftViewIdsArray, leftViewIdsArray_Count;

typedef struct MatchaStackScreenPBBar__storage_ {
  uint32_t _has_storage_[1];
  NSString *title;
  NSString *backButtonTitle;
  GPBInt64Array *rightViewIdsArray;
  GPBInt64Array *leftViewIdsArray;
  int64_t titleViewId;
} MatchaStackScreenPBBar__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "title",
        .dataTypeSpecific.className = NULL,
        .number = MatchaStackScreenPBBar_FieldNumber_Title,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaStackScreenPBBar__storage_, title),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "customBackButtonTitle",
        .dataTypeSpecific.className = NULL,
        .number = MatchaStackScreenPBBar_FieldNumber_CustomBackButtonTitle,
        .hasIndex = 3,
        .offset = 4,  // Stored in _has_storage_ to save space.
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeBool,
      },
      {
        .name = "backButtonTitle",
        .dataTypeSpecific.className = NULL,
        .number = MatchaStackScreenPBBar_FieldNumber_BackButtonTitle,
        .hasIndex = 5,
        .offset = (uint32_t)offsetof(MatchaStackScreenPBBar__storage_, backButtonTitle),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "titleViewId",
        .dataTypeSpecific.className = NULL,
        .number = MatchaStackScreenPBBar_FieldNumber_TitleViewId,
        .hasIndex = 6,
        .offset = (uint32_t)offsetof(MatchaStackScreenPBBar__storage_, titleViewId),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "rightViewIdsArray",
        .dataTypeSpecific.className = NULL,
        .number = MatchaStackScreenPBBar_FieldNumber_RightViewIdsArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaStackScreenPBBar__storage_, rightViewIdsArray),
        .flags = (GPBFieldFlags)(GPBFieldRepeated | GPBFieldPacked | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "leftViewIdsArray",
        .dataTypeSpecific.className = NULL,
        .number = MatchaStackScreenPBBar_FieldNumber_LeftViewIdsArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaStackScreenPBBar__storage_, leftViewIdsArray),
        .flags = (GPBFieldFlags)(GPBFieldRepeated | GPBFieldPacked | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "backButtonHidden",
        .dataTypeSpecific.className = NULL,
        .number = MatchaStackScreenPBBar_FieldNumber_BackButtonHidden,
        .hasIndex = 1,
        .offset = 2,  // Stored in _has_storage_ to save space.
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeBool,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaStackScreenPBBar class]
                                     rootClass:[MatchaStackScreenPBStacknavigatorRoot class]
                                          file:MatchaStackScreenPBStacknavigatorRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaStackScreenPBBar__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\006\002\025\000\003\017\000\004\013\000\005\000rightViewIds\000\006\000leftViewIds\000\007"
        "\020\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaStackScreenPBStackEvent

@implementation MatchaStackScreenPBStackEvent

@dynamic idArray, idArray_Count;

typedef struct MatchaStackScreenPBStackEvent__storage_ {
  uint32_t _has_storage_[1];
  GPBInt64Array *idArray;
} MatchaStackScreenPBStackEvent__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "idArray",
        .dataTypeSpecific.className = NULL,
        .number = MatchaStackScreenPBStackEvent_FieldNumber_IdArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaStackScreenPBStackEvent__storage_, idArray),
        .flags = (GPBFieldFlags)(GPBFieldRepeated | GPBFieldPacked),
        .dataType = GPBDataTypeInt64,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaStackScreenPBStackEvent class]
                                     rootClass:[MatchaStackScreenPBStacknavigatorRoot class]
                                          file:MatchaStackScreenPBStacknavigatorRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaStackScreenPBStackEvent__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
