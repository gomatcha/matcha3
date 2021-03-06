// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gomatcha.io/matcha/pb/view/imageview/imageview.proto

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

 #import "gomatcha.io/matcha/pb/view/imageview/Imageview.pbobjc.h"
 #import "gomatcha.io/matcha/pb/Color.pbobjc.h"
 #import "gomatcha.io/matcha/pb/Image.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - MatchaImageViewPBImageviewRoot

@implementation MatchaImageViewPBImageviewRoot

// No extensions in the file and none of the imports (direct or indirect)
// defined extensions, so no need to generate +extensionRegistry.

@end

#pragma mark - MatchaImageViewPBImageviewRoot_FileDescriptor

static GPBFileDescriptor *MatchaImageViewPBImageviewRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"matcha.view.imageview"
                                                 objcPrefix:@"MatchaImageViewPB"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - Enum MatchaImageViewPBResizeMode

GPBEnumDescriptor *MatchaImageViewPBResizeMode_EnumDescriptor(void) {
  static GPBEnumDescriptor *descriptor = NULL;
  if (!descriptor) {
    static const char *valueNames =
        "Fit\000Fill\000Stretch\000Center\000";
    static const int32_t values[] = {
        MatchaImageViewPBResizeMode_Fit,
        MatchaImageViewPBResizeMode_Fill,
        MatchaImageViewPBResizeMode_Stretch,
        MatchaImageViewPBResizeMode_Center,
    };
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(MatchaImageViewPBResizeMode)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:MatchaImageViewPBResizeMode_IsValidValue];
    if (!OSAtomicCompareAndSwapPtrBarrier(nil, worker, (void * volatile *)&descriptor)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL MatchaImageViewPBResizeMode_IsValidValue(int32_t value__) {
  switch (value__) {
    case MatchaImageViewPBResizeMode_Fit:
    case MatchaImageViewPBResizeMode_Fill:
    case MatchaImageViewPBResizeMode_Stretch:
    case MatchaImageViewPBResizeMode_Center:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - MatchaImageViewPBView

@implementation MatchaImageViewPBView

@dynamic hasImage, image;
@dynamic resizeMode;
@dynamic hasTint, tint;
@dynamic scale;

typedef struct MatchaImageViewPBView__storage_ {
  uint32_t _has_storage_[1];
  MatchaImageViewPBResizeMode resizeMode;
  MatchaPBImageOrResource *image;
  MatchaPBColor *tint;
  double scale;
} MatchaImageViewPBView__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "image",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBImageOrResource),
        .number = MatchaImageViewPBView_FieldNumber_Image,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaImageViewPBView__storage_, image),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "resizeMode",
        .dataTypeSpecific.enumDescFunc = MatchaImageViewPBResizeMode_EnumDescriptor,
        .number = MatchaImageViewPBView_FieldNumber_ResizeMode,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaImageViewPBView__storage_, resizeMode),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom | GPBFieldHasEnumDescriptor),
        .dataType = GPBDataTypeEnum,
      },
      {
        .name = "tint",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBColor),
        .number = MatchaImageViewPBView_FieldNumber_Tint,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(MatchaImageViewPBView__storage_, tint),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "scale",
        .dataTypeSpecific.className = NULL,
        .number = MatchaImageViewPBView_FieldNumber_Scale,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(MatchaImageViewPBView__storage_, scale),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeDouble,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaImageViewPBView class]
                                     rootClass:[MatchaImageViewPBImageviewRoot class]
                                          file:MatchaImageViewPBImageviewRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaImageViewPBView__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\001\002\n\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

int32_t MatchaImageViewPBView_ResizeMode_RawValue(MatchaImageViewPBView *message) {
  GPBDescriptor *descriptor = [MatchaImageViewPBView descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaImageViewPBView_FieldNumber_ResizeMode];
  return GPBGetMessageInt32Field(message, field);
}

void SetMatchaImageViewPBView_ResizeMode_RawValue(MatchaImageViewPBView *message, int32_t value) {
  GPBDescriptor *descriptor = [MatchaImageViewPBView descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:MatchaImageViewPBView_FieldNumber_ResizeMode];
  GPBSetInt32IvarWithFieldInternal(message, field, value, descriptor.file.syntax);
}


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
