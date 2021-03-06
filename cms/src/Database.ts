/*
* This file was generated by a tool.
* Rerun sql-ts to regenerate this file.
*/
export interface AdminEntity {
  'Email': string;
  'Id'?: number;
  'IsCalendarAdmin'?: boolean;
  'IsSuperAdmin'?: boolean;
  'IsTranslator'?: boolean;
  'LastUpdate': Date;
}
export interface AgeRatingEntity {
  'Code': string;
  'DescriptionId': number;
  'OrderIndex'?: number | null;
  'Title': string;
}
export interface AssetEntity {
  'AssetUuId': string;
  'Id'?: number;
}
export interface AudienceEntity {
  'Id'?: number;
  'LastUpdate': Date;
  'NameId': number;
}
export interface AudienceCalendarEventEntity {
  'AudienceId': number;
  'CalendarEventId': number;
}
export interface AudienceEpisodeEntity {
  'AudienceId': number;
  'EpisodeId': number;
}
export interface AudienceProgramEntity {
  'AudienceId': number;
  'ProgramId': number;
}
export interface AudienceSeasonEntity {
  'AudienceId': number;
  'SeasonId': number;
}
export interface AudienceSeriesEntity {
  'AudienceId': number;
  'SeriesId': number;
}
export interface AudienceTvGuideEntryEntity {
  'AudienceId': number;
  'TvGuideEntryId': number;
}
export interface AuditRecordEntity {
  'Details'?: string | null;
  'Event'?: string | null;
  'Id'?: number;
  'UserId'?: string | null;
}
export interface CalendarEventEntity {
  'AudienceId'?: number | null;
  'BccoEventId'?: string | null;
  'End'?: Date | null;
  'Id'?: number;
  'ImageUrl'?: string | null;
  'Locked'?: boolean | null;
  'NameId'?: number | null;
  'Start'?: Date | null;
  'UseImageFromBcco'?: boolean | null;
}
export interface CategoryEntity {
  'Id'?: number;
  'LastUpdate': Date;
  'NameId': number;
}
export interface CategoryEpisodeEntity {
  'CategoryId': number;
  'EpisodeId': number;
}
export interface CategoryProgramEntity {
  'CategoryId': number;
  'ProgramId': number;
}
export interface CategorySeasonEntity {
  'CategoryId': number;
  'SeasonId': number;
}
export interface CategorySeriesEntity {
  'CategoryId': number;
  'SeriesId': number;
}
export interface DownloadableEntity {
  'BlobContainerName': string;
  'Filename': string;
  'GroupName': string;
  'Id'?: number;
  'LanguageId'?: number | null;
  'LanguageThreeLetterCode'?: string | null;
  'SizeInBytes'?: string | null;
}
export interface EpisodeEntity {
  'AllowSpecialAccess'?: boolean;
  'AllowSpecialAccessFKTB'?: boolean;
  'AvailableFrom'?: Date | null;
  'AvailableIn'?: string | null;
  'AvailableTo'?: Date | null;
  'DescriptionId': number;
  'EpisodeNo': number;
  'Id'?: number;
  'Image'?: string | null;
  'LastUpdate': Date;
  'LongDescriptionId'?: number | null;
  'OldGoodie'?: boolean;
  'Published': Date;
  'Recommended'?: boolean;
  'SearchId'?: number | null;
  'SeasonId': number;
  'Status': number;
  'TitleId': number;
  'VideoId': number;
  'ViewCount'?: number;
  'Visibility'?: number;
}
export interface EpisodeDownloadableEntity {
  'DownloadableId': number;
  'EpisodeId': number;
}
export interface FaqCategoryEntity {
  'Id'?: number;
  'LastUpdate'?: Date;
  'OrderIndex': number;
  'Status'?: number;
  'TitleId': number;
  'Visibility'?: number;
}
export interface FaqCategoryFaqItemEntity {
  'FaqCategoryId': number;
  'FaqItemId': number;
  'OrderIndex': number;
  'Status'?: number;
  'Visibility'?: number;
}
export interface FaqItemEntity {
  'AnswerId': number;
  'Id'?: number;
  'LastUpdate'?: Date;
  'QuestionId': number;
  'Status'?: number;
}
export interface FeaturedEntity {
  'AppImageUrl'?: string | null;
  'CategoryId'?: number | null;
  'DescriptionId': number;
  'EpisodeId'?: number | null;
  'Id'?: number;
  'IsLive'?: boolean;
  'LastUpdate'?: Date | null;
  'Order'?: number | null;
  'ProgramId'?: number | null;
  'SeasonId'?: number | null;
  'SeriesId'?: number | null;
  'Status': number;
  'TitleId': number;
  'TvImageUrl'?: string | null;
  'Type': number;
  'Url': string;
  'Visibility'?: number;
}
export interface JobEntity {
  'CreateDateTime': Date;
  'Id'?: number;
  'JobUuId': string;
  'VideoId': number;
}
export interface LanguageEntity {
  'CultureCode': string;
  'Id'?: number;
  'Name': string;
}
export interface LocalizableStringEntity {
  'Id'?: number;
  'Key'?: string | null;
}
export interface LocalizedStringEntity {
  'Id'?: number;
  'LanguageId': number;
  'ParentId': number;
  'Value': string;
}
export interface MetadataEntity {
  'Id'?: number;
  'MetaValue': string;
  'Name': string;
  'VideoId': number;
}
export interface NotificationEntity {
  'Badge'?: number | null;
  'Command'?: string | null;
  'ExternalId'?: string | null;
  'Id'?: number;
  'LastModified': Date;
  'MessageId'?: number | null;
  'SendDateTime'?: Date | null;
  'Status'?: number;
  'Tag'?: string | null;
  'TitleId'?: number | null;
}
export interface ProgramEntity {
  'AgeRatingCode'?: string | null;
  'AllowSpecialAccess'?: boolean;
  'AllowSpecialAccessFKTB'?: boolean;
  'AvailableFrom'?: Date | null;
  'AvailableIn'?: string | null;
  'AvailableTo'?: Date | null;
  'DescriptionId': number;
  'Id'?: number;
  'Image'?: string | null;
  'LastUpdate': Date;
  'LongDescriptionId'?: number | null;
  'OldGoodie'?: boolean;
  'Published': Date;
  'Recommended'?: boolean;
  'SearchId'?: number | null;
  'Status': number;
  'TitleId': number;
  'VideoId': number;
  'ViewCount'?: number;
  'Visibility'?: number;
}
export interface SeasonEntity {
  'AgeRatingCode'?: string | null;
  'AvailableFrom'?: Date | null;
  'AvailableIn'?: string | null;
  'AvailableTo'?: Date | null;
  'DescriptionId': number;
  'Id'?: number;
  'Image'?: string | null;
  'LastUpdate': Date;
  'OldGoodie'?: boolean;
  'Published': Date;
  'Recommended'?: boolean;
  'SeasonNo': number;
  'SeriesId': number;
  'Status': number;
  'TitleId': number;
}
export interface SeriesEntity {
  'AgeRatingCode'?: string | null;
  'AvailableFrom'?: Date | null;
  'AvailableIn'?: string | null;
  'AvailableTo'?: Date | null;
  'DescriptionId': number;
  'Id'?: number;
  'Image'?: string | null;
  'IsCategory'?: number;
  'LastUpdate': Date;
  'OldGoodie'?: boolean;
  'Published': Date;
  'Recommended'?: boolean;
  'SearchId'?: number | null;
  'ShowAsBundle'?: number;
  'ShowEpisodeTitles'?: number;
  'Status': number;
  'TitleId': number;
}
export interface SliderEntity {
  'ExplorerOptions': string;
  'GetEpisodes'?: boolean;
  'Id'?: number;
  'LastUpdate'?: Date;
  'OrderIndex': number;
  'Status'?: number;
  'TitleId': number;
  'Visibility'?: number;
}
export interface SystemDataEntity {
  'Id'?: number;
  'Key': string;
  'Value': string;
}
export interface TaskEntity {
  'Id'?: number;
  'JobId': number;
  'TaskUuId': string;
}
export interface TaskAssetEntity {
  'AssetId': number;
  'IsInputAsset': boolean;
  'TaskId': number;
}
export interface TvGuideEntryEntity {
  'AudienceId'?: number | null;
  'DescriptionId'?: number | null;
  'End'?: Date | null;
  'EventId'?: number | null;
  'Id'?: number;
  'ImageUrl'?: string | null;
  'Start'?: Date | null;
  'Status'?: number;
  'TitleId'?: number | null;
  'UseImageFromVod'?: boolean | null;
  'VodId'?: number | null;
  'VodType'?: string | null;
}
export interface VideoEntity {
  'EncodingStatus'?: number | null;
  'EncodingVersion'?: number | null;
  'Filename': string;
  'Id'?: number;
  'IsEncrypted'?: boolean;
  'LastUpdate'?: Date | null;
  'Length'?: string | null;
}
export interface VideoThumbnailEntity {
  'Id'?: number;
  'ThumbnailUrl': string;
  'VideoId': number;
}
export interface VideoUrlEntity {
  'EncodingType': string;
  'Id'?: number;
  'IsVideoClip': boolean;
  'VideoId': number;
  'VideoUrl': string;
}
