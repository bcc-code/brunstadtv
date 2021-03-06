import { defineHook } from '@directus/extensions-sdk';
import { createEpisodeTranslation, createSeasonTranslation, createShowTranslation, updateEpisodeTranslation, updateSeasonTranslation, updateShowTranslation } from './filters/translations';
import { createShow, deleteShow, updateShow } from './filters/shows';
import { createEpisode, deleteEpisode, updateEpisode } from './filters/episodes';
import { createSeason, deleteSeason, updateSeason } from './filters/seasons';
import { createEpisodesUsergroup, deleteEpisodesUsergroup, deleteEpisodesUsergroupEarlyAccess, createEpisodesUsergroupEarlyAccess, updateUsergroup } from './filters/usergroups';
import { createList, deleteList, updateList } from './filters/lists';
import { createListRelation, deleteListRelation } from './filters/lists_relations';
import { createAsset, deleteAsset, updateAsset } from './filters/assets';
import { createAssetstream, deleteAssetstream, updateAssetstream } from './filters/assetstreams';
import { createEpisodeTag, deleteEpisodeTag, updateTag } from './filters/tags';


export default defineHook(({ filter }, {}) => {
	if (process.env.LEGACY_SYNC === "off") {
		return
	}

	filter('items.create', createShow)
	filter('items.create', createSeason)
	filter('items.create', createEpisode)
	filter('items.create', createList)
	filter('items.create', createShowTranslation);
	filter('items.create', createSeasonTranslation);
	filter('items.create', createEpisodeTranslation);
	filter('items.create', createEpisodesUsergroup);
	filter('items.create', createEpisodesUsergroupEarlyAccess);
	filter('items.create', createListRelation);
	filter('items.create', createAsset);
	filter('items.create', createAssetstream);
	filter('items.create', createEpisodeTag);

	filter('items.update', updateShow)
	filter('items.update', updateSeason)
	filter('items.update', updateEpisode)
	filter('items.update', updateList)
	filter('items.update', updateShowTranslation)
	filter('items.update', updateSeasonTranslation)
	filter('items.update', updateEpisodeTranslation)
	filter('items.update', updateUsergroup)
	filter('items.update', updateAsset)
	filter('items.update', updateAssetstream)
	filter('items.update', updateTag)

	filter('items.delete', deleteEpisodesUsergroup);
	filter('items.delete', deleteEpisodesUsergroupEarlyAccess);
	filter('items.delete', deleteEpisode);
	filter('items.delete', deleteSeason);
	filter('items.delete', deleteShow);
	filter('items.delete', deleteList);
	filter('items.delete', deleteListRelation);
	filter('items.delete', deleteAsset);
	filter('items.delete', deleteAssetstream);
	filter('items.delete', deleteEpisodeTag);

});

