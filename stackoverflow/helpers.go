package stackoverflow

import (
	so "terraform-provider-stackoverflow/stackoverflow/client"
)

func selectTagNamesToArray(tags []so.Tag) []string {
	tagCollection := make([]string, len(tags))

	for i, tag := range tags {
		tagCollection[i] = tag.Name
	}

	return tagCollection
}

func selectCollectionContentIdsToArray(content []so.CollectionContent) []int {
	contentCollection := make([]int, len(content))

	for i, item := range content {
		contentCollection[i] = item.ID
	}

	return contentCollection
}

func convertToArray[T any](items []interface{}) []T {
	collection := make([]T, len(items))

	for i, item := range items {
		collection[i] = item.(T)
	}

	return collection
}

func mergeDefaultTagsWithResourceTags(defaultTags []string, tags []string) []string {
	allTags := append(defaultTags, tags...)
	tagMap := map[string]bool{}
	tagCollection := []string{}

	for _, tag := range allTags {
		if !tagMap[tag] {
			tagMap[tag] = true
			tagCollection = append(tagCollection, tag)
		}
	}

	return tagCollection
}

func ignoreDefaultTags(defaultTags []string, actualTags []string, resourceTags []string) []string {
	tagMap := map[string]bool{}
	resourceTagMap := map[string]bool{}
	tagCollection := []string{}

	for _, tag := range actualTags {
		if !tagMap[tag] {
			tagMap[tag] = true
		}
	}

	for _, tag := range resourceTags {
		if !resourceTagMap[tag] {
			resourceTagMap[tag] = true
		}
	}

	for _, tag := range defaultTags {
		if tagMap[tag] && !resourceTagMap[tag] {
			tagMap[tag] = false
		}
	}

	for key := range tagMap {
		if tagMap[key] {
			tagCollection = append(tagCollection, key)
		}
	}

	return tagCollection
}
