package stackoverflow

func expandTagsToArray(tags []interface{}) []string {
	tagCollection := make([]string, len(tags))

	for i, tag := range tags {
		tagCollection[i] = tag.(string)
	}

	return tagCollection
}
