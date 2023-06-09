package stackoverflow

import (
	"context"
	"fmt"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFilter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFilterCreate,
		ReadContext:   resourceFilterRead,
		//UpdateContext: resourceFilterUpdate,
		DeleteContext: resourceFilterDelete,
		Schema:        map[string]*schema.Schema{},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	filter := &so.Filter{
		Include: getFilterIncludes(),
		Exclude: getFilterExcludes(),
		Unsafe:  true,
	}

	newFilter, err := client.CreateFilter(filter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(newFilter.ID)

	return diags
}

func resourceFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)
	var diags diag.Diagnostics
	filterIDs := []string{d.Id()}
	filters, err := client.GetFilters(&filterIDs)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(*filters) < 1 {
		return diag.FromErr(fmt.Errorf("no filter found matching identifier %s", d.Id()))
	}

	if len(*filters) > 1 {
		return diag.FromErr(fmt.Errorf("found %d filters matching identifier %s", len(*filters), d.Id()))
	}

	filter := (*filters)[0]

	d.SetId(filter.ID)

	return diags
}

//func resourceFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
//	client := meta.(*so.Client)
//
//	// Warning or errors can be collected in a slice type
//	var diags diag.Diagnostics
//
//	filter := &so.Filter{
//		ID: d.Id(),
//	}
//
//	_, err2 := client.UpdateFilter(filter)
//	if err2 != nil {
//		return diag.FromErr(err2)
//	}
//
//	return diags
//}

func resourceFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)
	var diags diag.Diagnostics

	err2 := client.DeleteFilter(d.Id())

	if err2 != nil {
		return diag.FromErr(err2)
	}

	return diags
}

func getFilterIncludes() []string {
	includes := []string{
		".backoff",
		".error_id",
		".error_message",
		".error_name",
		".has_more",
		".items",
		".page",
		".page_size",
		".quota_max",
		".quota_remaining",
		".total",
		".type",
		"access_token.access_token",
		"access_token.account_id",
		"access_token.expires_on_date",
		"access_token.scope",
		"account_merge.merge_date",
		"account_merge.new_account_id",
		"account_merge.old_account_id",
		"achievement.account_id",
		"achievement.achievement_type",
		"achievement.badge_rank",
		"achievement.creation_date",
		"achievement.is_unread",
		"achievement.link",
		"achievement.on_site",
		"achievement.reputation_change",
		"achievement.title",
		"answer.accepted",
		"answer.answer_id",
		"answer.body",
		"answer.body_markdown",
		"answer.can_comment",
		"answer.can_edit",
		"answer.can_flag",
		"answer.can_suggest_edit",
		"answer.collectives",
		"answer.community_owned_date",
		"answer.content_license",
		"answer.creation_date",
		"answer.is_accepted",
		"answer.last_activity_date",
		"answer.last_edit_date",
		"answer.last_editor",
		"answer.locked_date",
		"answer.owner",
		"answer.posted_by_collectives",
		"answer.question_id",
		"answer.recommendations",
		"answer.score",
		"answer.tags",
		"answer.title",
		"article.article_id",
		"article.article_type",
		"article.body",
		"article.body_markdown",
		"article.creation_date",
		"article.last_activity_date",
		"article.last_edit_date",
		"article.last_editor",
		"article.link",
		"article.owner",
		"article.score",
		"article.tags",
		"article.title",
		"article.view_count",
		"badge.award_count",
		"badge.badge_id",
		"badge.badge_type",
		"badge.link",
		"badge.name",
		"badge.rank",
		"badge.user",
		"closed_details.by_users",
		"closed_details.description",
		"closed_details.on_hold",
		"closed_details.original_questions",
		"closed_details.reason",
		"collective.description",
		"collective.external_links",
		"collective.link",
		"collective.name",
		"collective.slug",
		"collective.tags",
		"collective_external_link.link",
		"collective_external_link.type",
		"collective_membership.collective",
		"collective_membership.role",
		"collective_recommendation.collective",
		"collective_recommendation.creation_date",
		"collective_report.collective",
		"collective_report.creation_date",
		"collective_report.dimensions",
		"collective_report.download_link",
		"collective_report.end_date",
		"collective_report.included_tags",
		"collective_report.metrics",
		"collective_report.name",
		"collective_report.report_id",
		"collective_report.report_type",
		"collective_report.start_date",
		"collective_report.state",
		"comment.comment_id",
		"comment.content_license",
		"comment.creation_date",
		"comment.edited",
		"comment.owner",
		"comment.post_id",
		"comment.reply_to_user",
		"comment.score",
		"error.description",
		"error.error_id",
		"error.error_name",
		"event.creation_date",
		"event.event_id",
		"event.event_type",
		"exchanged_access_token.access_token",
		"exchanged_access_token.account_id",
		"exchanged_access_token.exchange_type",
		"exchanged_access_token.expires_on_date",
		"exchanged_access_token.original_access_token",
		"exchanged_access_token.scope",
		"filter.filter",
		"filter.filter_type",
		"filter.included_fields",
		"flag_option.count",
		"flag_option.description",
		"flag_option.dialog_title",
		"flag_option.has_flagged",
		"flag_option.is_retraction",
		"flag_option.option_id",
		"flag_option.requires_comment",
		"flag_option.requires_question_id",
		"flag_option.requires_site",
		"flag_option.sub_options",
		"flag_option.title",
		"inbox_item.answer_id",
		"inbox_item.comment_id",
		"inbox_item.creation_date",
		"inbox_item.is_unread",
		"inbox_item.item_type",
		"inbox_item.link",
		"inbox_item.question_id",
		"inbox_item.site",
		"inbox_item.title",
		"info.answers_per_minute",
		"info.api_revision",
		"info.badges_per_minute",
		"info.new_active_users",
		"info.questions_per_minute",
		"info.total_accepted",
		"info.total_answers",
		"info.total_badges",
		"info.total_comments",
		"info.total_questions",
		"info.total_unanswered",
		"info.total_users",
		"info.total_votes",
		"migration_info.on_date",
		"migration_info.other_site",
		"migration_info.question_id",
		"network_activity.account_id",
		"network_activity.activity_type",
		"network_activity.api_site_parameter",
		"network_activity.badge_id",
		"network_activity.creation_date",
		"network_activity.description",
		"network_activity.link",
		"network_activity.post_id",
		"network_activity.score",
		"network_activity.tags",
		"network_activity.title",
		"network_post.post_id",
		"network_post.post_type",
		"network_post.score",
		"network_post.title",
		"network_user.account_id",
		"network_user.answer_count",
		"network_user.badge_counts",
		"network_user.creation_date",
		"network_user.last_access_date",
		"network_user.question_count",
		"network_user.reputation",
		"network_user.site_name",
		"network_user.site_url",
		"network_user.user_id",
		"notice.body",
		"notice.creation_date",
		"notice.owner_user_id",
		"notification.body",
		"notification.creation_date",
		"notification.is_unread",
		"notification.notification_type",
		"notification.post_id",
		"notification.site",
		"original_question.accepted_answer_id",
		"original_question.answer_count",
		"original_question.question_id",
		"original_question.title",
		"post.collectives",
		"post.content_license",
		"post.creation_date",
		"post.last_activity_date",
		"post.last_edit_date",
		"post.link",
		"post.owner",
		"post.post_id",
		"post.post_type",
		"post.posted_by_collectives",
		"post.score",
		"privilege.description",
		"privilege.reputation",
		"privilege.short_description",
		"question.accepted_answer_id",
		"question.answer_count",
		"question.body",
		"question.body_markdown",
		"question.bounty_amount",
		"question.bounty_closes_date",
		"question.can_answer",
		"question.can_close",
		"question.can_comment",
		"question.can_edit",
		"question.can_flag",
		"question.can_suggest_edit",
		"question.closed_date",
		"question.closed_reason",
		"question.collectives",
		"question.community_owned_date",
		"question.content_license",
		"question.creation_date",
		"question.is_answered",
		"question.last_activity_date",
		"question.last_edit_date",
		"question.last_editor",
		"question.link",
		"question.locked_date",
		"question.migrated_from",
		"question.migrated_to",
		"question.owner",
		"question.posted_by_collectives",
		"question.protected_date",
		"question.question_id",
		"question.score",
		"question.tags",
		"question.title",
		"question.view_count",
		"question_timeline.comment_id",
		"question_timeline.content_license",
		"question_timeline.creation_date",
		"question_timeline.down_vote_count",
		"question_timeline.owner",
		"question_timeline.post_id",
		"question_timeline.question_id",
		"question_timeline.revision_guid",
		"question_timeline.timeline_type",
		"question_timeline.up_vote_count",
		"question_timeline.user",
		"related_site.api_site_parameter",
		"related_site.name",
		"related_site.relation",
		"related_site.site_url",
		"reputation.on_date",
		"reputation.post_id",
		"reputation.post_type",
		"reputation.reputation_change",
		"reputation.user_id",
		"reputation.vote_type",
		"reputation_history.creation_date",
		"reputation_history.post_id",
		"reputation_history.reputation_change",
		"reputation_history.reputation_history_type",
		"reputation_history.user_id",
		"revision.comment",
		"revision.content_license",
		"revision.creation_date",
		"revision.is_rollback",
		"revision.last_tags",
		"revision.last_title",
		"revision.post_id",
		"revision.post_type",
		"revision.revision_guid",
		"revision.revision_number",
		"revision.revision_type",
		"revision.set_community_wiki",
		"revision.tags",
		"revision.title",
		"revision.user",
		"search_excerpt.answer_count",
		"search_excerpt.answer_id",
		"search_excerpt.body",
		"search_excerpt.creation_date",
		"search_excerpt.equivalent_tag_search",
		"search_excerpt.excerpt",
		"search_excerpt.has_accepted_answer",
		"search_excerpt.is_accepted",
		"search_excerpt.is_answered",
		"search_excerpt.item_type",
		"search_excerpt.last_activity_date",
		"search_excerpt.question_id",
		"search_excerpt.question_score",
		"search_excerpt.score",
		"search_excerpt.tags",
		"search_excerpt.title",
		"shallow_user.accept_rate",
		"shallow_user.account_id",
		"shallow_user.display_name",
		"shallow_user.link",
		"shallow_user.profile_image",
		"shallow_user.reputation",
		"shallow_user.user_id",
		"shallow_user.user_type",
		"site.aliases",
		"site.api_site_parameter",
		"site.audience",
		"site.closed_beta_date",
		"site.favicon_url",
		"site.high_resolution_icon_url",
		"site.icon_url",
		"site.launch_date",
		"site.logo_url",
		"site.markdown_extensions",
		"site.name",
		"site.open_beta_date",
		"site.related_sites",
		"site.site_state",
		"site.site_type",
		"site.site_url",
		"site.styling",
		"site.twitter_account",
		"styling.link_color",
		"styling.tag_background_color",
		"styling.tag_foreground_color",
		"suggested_edit.approval_date",
		"suggested_edit.comment",
		"suggested_edit.creation_date",
		"suggested_edit.post_id",
		"suggested_edit.post_type",
		"suggested_edit.proposing_user",
		"suggested_edit.rejection_date",
		"suggested_edit.suggested_edit_id",
		"suggested_edit.tags",
		"suggested_edit.title",
		"tag.collectives",
		"tag.count",
		"tag.has_synonyms",
		"tag.is_moderator_only",
		"tag.is_required",
		"tag.name",
		"tag.user_id",
		"tag_preference.tag_name",
		"tag_preference.tag_preference_type",
		"tag_preference.user_id",
		"tag_score.post_count",
		"tag_score.score",
		"tag_score.user",
		"tag_synonym.applied_count",
		"tag_synonym.creation_date",
		"tag_synonym.from_tag",
		"tag_synonym.last_applied_date",
		"tag_synonym.to_tag",
		"tag_wiki.body_last_edit_date",
		"tag_wiki.excerpt",
		"tag_wiki.excerpt_last_edit_date",
		"tag_wiki.tag_name",
		"top_tag.answer_count",
		"top_tag.answer_score",
		"top_tag.question_count",
		"top_tag.question_score",
		"top_tag.tag_name",
		"top_tag.user_id",
		"user.accept_rate",
		"user.account_id",
		"user.age",
		"user.badge_counts",
		"user.collectives",
		"user.creation_date",
		"user.display_name",
		"user.is_employee",
		"user.last_access_date",
		"user.last_modified_date",
		"user.link",
		"user.location",
		"user.profile_image",
		"user.reputation",
		"user.reputation_change_day",
		"user.reputation_change_month",
		"user.reputation_change_quarter",
		"user.reputation_change_week",
		"user.reputation_change_year",
		"user.timed_penalty_date",
		"user.user_id",
		"user.user_type",
		"user.website_url",
		"user_timeline.badge_id",
		"user_timeline.comment_id",
		"user_timeline.creation_date",
		"user_timeline.detail",
		"user_timeline.post_id",
		"user_timeline.post_type",
		"user_timeline.suggested_edit_id",
		"user_timeline.timeline_type",
		"user_timeline.title",
		"user_timeline.user_id",
		"write_permission.can_add",
		"write_permission.can_delete",
		"write_permission.can_edit",
		"write_permission.max_daily_actions",
		"write_permission.min_seconds_between_actions",
		"write_permission.object_type",
		"write_permission.user_id",
	}

	return includes
}

func getFilterExcludes() []string {
	excludes := []string{
		"article.comment_count",
		"article.comments",
		"article.last_editor",
		"comment.body",
		"comment.body_markdown",
		"comment.can_flag",
		"comment.link",
		"comment.post_type",
		"comment.upvoted",
		"shallow_user.badge_counts",
		"answer.awarded_bounty_amount",
		"answer.awarded_bounty_users",
		"answer.comment_count",
		"answer.comments",
		"answer.down_vote_count",
		"answer.downvoted",
		"answer.last_editor",
		"answer.link",
		"answer.share_link",
		"answer.up_vote_count",
		"answer.upvoted",
		"badge.description",
		"event.excerpt",
		"event.link",
		"flag_option.question",
		"inbox_item.body",
		"info.site",
		"network_user.top_answers",
		"network_user.top_questions",
		"network_user.user_type",
		"post.body",
		"post.body_markdown",
		"post.comment_count",
		"post.comments",
		"post.down_vote_count",
		"post.downvoted",
		"post.last_editor",
		"post.share_link",
		"post.title",
		"post.up_vote_count",
		"post.upvoted",
		"question.answers",
		"question.bounty_user",
		"question.close_vote_count",
		"question.closed_details",
		"question.comment_count",
		"question.comments",
		"question.delete_vote_count",
		"question.down_vote_count",
		"question.downvoted",
		"question.favorite_count",
		"question.favorited",
		"question.last_editor",
		"question.notice",
		"question.reopen_vote_count",
		"question.share_link",
		"question.up_vote_count",
		"question.upvoted",
		"reputation.link",
		"reputation.title",
		"revision.body",
		"revision.last_body",
		"search_excerpt.closed_date",
		"search_excerpt.community_owned_date",
		"search_excerpt.last_activity_user",
		"search_excerpt.locked_date",
		"search_excerpt.owner",
		"suggested_edit.body",
		"tag.last_activity_date",
		"tag.synonyms",
		"tag_wiki.body",
		"tag_wiki.last_body_editor",
		"tag_wiki.last_excerpt_editor",
		"user.about_me",
		"user.answer_count",
		"user.down_vote_count",
		"user.question_count",
		"user.up_vote_count",
		"user.view_count",
		"user_timeline.link",
	}

	return excludes
}
