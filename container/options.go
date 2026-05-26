// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package container

// options.go defines container options.

import (
	"image"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgetapi"
)

// applyOptions applies the options to the container and validates them.
func applyOptions(c *Container, opts ...Option) error { _ = "STUB: not implemented"; return nil }

// ensure all the container identifiers are either empty or unique.
func validateIds(c *Container, seen map[string]bool) error { _ = "STUB: not implemented"; return nil }

// ensure all the container only have one split modifier.
func validateSplits(c *Container) error { _ = "STUB: not implemented"; return nil }

// validateOptions validates options set in the container tree.
func validateOptions(c *Container) error { _ = "STUB: not implemented"; return nil }

// Option is used to provide options to a container.
type Option interface {
	// set sets the provided option.
	set(*Container) error
}

// options stores the options provided to the container.
type options struct {
	// id is the identifier provided by the user.
	id string

	// global are options that apply globally to all containers in the tree.
	// There is only one instance of these options in the entire tree, if any
	// of the child containers change their values, the new values apply to the
	// entire container tree.
	global *globalOptions

	// inherited are options that are inherited by child containers.
	// After inheriting these options, the child container can set them to
	// different values.
	inherited inherited

	// split identifies how is this container split.
	split         splitType
	splitReversed bool
	splitPercent  int
	splitFixed    int

	// widget is the widget in the container.
	// A container can have either two sub containers (left and right) or a
	// widget. But not both.
	widget widgetapi.Widget

	// Alignment of the widget if present.
	hAlign align.Horizontal
	vAlign align.Vertical

	// border is the border around the container.
	border            linestyle.LineStyle
	borderTitle       string
	borderTitleHAlign align.Horizontal

	// padding is a space reserved between the outer edge of the container and
	// its content (the widget or other sub-containers).
	padding padding

	// margin is a space reserved on the outside of the container.
	margin margin

	// keyFocusSkip asserts whether this container should be skipped when focus
	// is being moved using either of KeyFocusNext or KeyFocusPrevious.
	keyFocusSkip bool
	// keyFocusGroups are the focus groups this container belongs to.
	keyFocusGroups []FocusGroup
}

// margin stores the configured margin for the container.
// For each margin direction, only one of the percentage or cells is set.
type margin struct {
	topCells    int
	topPerc     int
	rightCells  int
	rightPerc   int
	bottomCells int
	bottomPerc  int
	leftCells   int
	leftPerc    int
}

// apply applies the configured margin to the area.
func (p *margin) apply(ar image.Rectangle) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), nil
}

// padding stores the configured padding for the container.
// For each padding direction, only one of the percentage or cells is set.
type padding struct {
	topCells    int
	topPerc     int
	rightCells  int
	rightPerc   int
	bottomCells int
	bottomPerc  int
	leftCells   int
	leftPerc    int
}

// apply applies the configured padding to the area.
func (p *padding) apply(ar image.Rectangle) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), nil
}

// inherited contains options that are inherited by child containers.
type inherited struct {
	// borderColor is the color used for the border.
	borderColor cell.Color
	// focusedColor is the color used for the border when focused.
	focusedColor cell.Color
	// titleColor is the color used for the title.
	titleColor *cell.Color
	// titleFocusedColor is the color used for the title when focused.
	titleFocusedColor *cell.Color
}

// focusGroups maps focus group numbers that have the same key assigned.
// The value is always true for all the keys.
type focusGroups map[FocusGroup]bool

// firstMatching examines the focus groups the container is assigned to and
// returns the first matching focus group that is also present in this
// instance. The bool return value indicates if match was found.
func (fg focusGroups) firstMatching(contGroups []FocusGroup) (bool, FocusGroup) {
	_ = "STUB: not implemented"
	return false, *new(FocusGroup)
}

// globalOptions are options that can only have a single value across the
// entire tree of containers.
// Regardless of which container they get set on, the new value will take
// effect on all the containers in the tree.
type globalOptions struct {
	// keyFocusNext when set is the key that moves the focus to the next container.
	keyFocusNext *keyboard.Key
	// keyFocusPrevious when set is the key that moves the focus to the previous container.
	keyFocusPrevious *keyboard.Key
	// keysFocusGroupNext maps keyboard keys that move to the next container
	// within a focus group to the focus groups they should work on in the
	// order they were configured.
	keyFocusGroupsNext map[keyboard.Key]focusGroups
	// keysFocusGroupPrevious maps keyboard keys that move to the previous
	// container within a focus group to the focus groups they should work on
	// in the order they were configured.
	keyFocusGroupsPrevious map[keyboard.Key]focusGroups
}

// newOptions returns a new options instance with the default values.
// Parent are the inherited options from the parent container or nil if these
// options are for a container with no parent (the root).
func newOptions(parent *options) *options { _ = "STUB: not implemented"; return nil }

// option implements Option.
type option func(*Container) error

// set implements Option.set.
func (o option) set(c *Container) error {
	_ = "STUB: not implemented"

	// SplitOption is used when splitting containers.
	return nil
}

type SplitOption interface {
	// setSplit sets the provided split option.
	setSplit(*options) error
}

// splitOption implements SplitOption.
type splitOption func(*options) error

// setSplit implements SplitOption.setSplit.
func (so splitOption) setSplit(opts *options) error {
	_ = "STUB: not implemented"

	// DefaultSplitReversed is the default value for the SplitReversed option.
	return nil
}

const DefaultSplitReversed = false

// DefaultSplitPercent is the default value for the SplitPercent option.
const DefaultSplitPercent = 50

// DefaultSplitFixed is the default value for the SplitFixed option.
const DefaultSplitFixed = -1

// SplitPercent sets the relative size of the split as percentage of the
// available space.
// When using SplitVertical, the provided size is applied to the new left
// container, the new right container gets the reminder of the size.
// When using SplitHorizontal, the provided size is applied to the new top
// container, the new bottom container gets the reminder of the size.
// The provided value must be a positive number in the range 0 < p < 100.
// If not provided, defaults to DefaultSplitPercent.
func SplitPercent(p int) SplitOption { _ = "STUB: not implemented"; return *new(SplitOption) }

// SplitPercentFromEnd sets the relative size of the split as percentage of the
// available space.
// When using SplitVertical, the provided size is applied to the new right
// container, the new left container gets the reminder of the size.
// When using SplitHorizontal, the provided size is applied to the new bottom
// container, the new top container gets the reminder of the size.
// The provided value must be a positive number in the range 0 < p < 100.
// If not provided, defaults to using SplitPercent with DefaultSplitPercent.
func SplitPercentFromEnd(p int) SplitOption { _ = "STUB: not implemented"; return *new(SplitOption) }

// SplitFixed sets the size of the first container to be a fixed value
// and makes the second container take up the remaining space.
// When using SplitVertical, the provided size is applied to the new left
// container, the new right container gets the reminder of the size.
// When using SplitHorizontal, the provided size is applied to the new top
// container, the new bottom container gets the reminder of the size.
// The provided value must be a positive number in the range 0 <= cells.
// If SplitFixed* or SplitPercent* is not specified, it defaults to
// SplitPercent() and its given value.
// Only one SplitFixed* or SplitPercent* may be specified per container.
func SplitFixed(cells int) SplitOption { _ = "STUB: not implemented"; return *new(SplitOption) }

// SplitFixedFromEnd sets the size of the second container to be a fixed value
// and makes the first container take up the remaining space.
// When using SplitVertical, the provided size is applied to the new right
// container, the new left container gets the reminder of the size.
// When using SplitHorizontal, the provided size is applied to the new bottom
// container, the new top container gets the reminder of the size.
// The provided value must be a positive number in the range 0 <= cells.
// If SplitFixed* or SplitPercent* is not specified, it defaults to
// SplitPercent() and its given value.
// Only one SplitFixed* or SplitPercent* may be specified per container.
func SplitFixedFromEnd(cells int) SplitOption { _ = "STUB: not implemented"; return *new(SplitOption) }

// SplitVertical splits the container along the vertical axis into two sub
// containers. The use of this option removes any widget placed at this
// container, containers with sub containers cannot contain widgets.
func SplitVertical(l LeftOption, r RightOption, opts ...SplitOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// SplitHorizontal splits the container along the horizontal axis into two sub
// containers. The use of this option removes any widget placed at this
// container, containers with sub containers cannot contain widgets.
func SplitHorizontal(t TopOption, b BottomOption, opts ...SplitOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// ID sets an identifier for this container.
// This ID can be later used to perform dynamic layout changes by passing new
// options to this container. When provided, it must be a non-empty string that
// is unique among all the containers.
func ID(id string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Clear clears this container.
// If the container contains a widget, the widget is removed.
// If the container had any sub containers or splits, they are removed.
func Clear() Option { _ = "STUB: not implemented"; return *new(Option) }

// PlaceWidget places the provided widget into the container.
// The use of this option removes any sub containers. Containers with sub
// containers cannot have widgets.
func PlaceWidget(w widgetapi.Widget) Option { _ = "STUB: not implemented"; return *new(Option) }

// MarginTop sets reserved space outside of the container at its top.
// The provided number is the absolute margin in cells and must be zero or a
// positive integer. Only one of MarginTop or MarginTopPercent can be specified.
func MarginTop(cells int) Option { _ = "STUB: not implemented"; return *new(Option) }

// MarginRight sets reserved space outside of the container at its right.
// The provided number is the absolute margin in cells and must be zero or a
// positive integer. Only one of MarginRight or MarginRightPercent can be specified.
func MarginRight(cells int) Option { _ = "STUB: not implemented"; return *new(Option) }

// MarginBottom sets reserved space outside of the container at its bottom.
// The provided number is the absolute margin in cells and must be zero or a
// positive integer. Only one of MarginBottom or MarginBottomPercent can be specified.
func MarginBottom(cells int) Option { _ = "STUB: not implemented"; return *new(Option) }

// MarginLeft sets reserved space outside of the container at its left.
// The provided number is the absolute margin in cells and must be zero or a
// positive integer. Only one of MarginLeft or MarginLeftPercent can be specified.
func MarginLeft(cells int) Option { _ = "STUB: not implemented"; return *new(Option) }

// MarginTopPercent sets reserved space outside of the container at its top.
// The provided number is a relative margin defined as percentage of the container's height.
// Only one of MarginTop or MarginTopPercent can be specified.
// The value must be in range 0 <= value <= 100.
func MarginTopPercent(perc int) Option { _ = "STUB: not implemented"; return *new(Option) }

// MarginRightPercent sets reserved space outside of the container at its right.
// The provided number is a relative margin defined as percentage of the container's height.
// Only one of MarginRight or MarginRightPercent can be specified.
// The value must be in range 0 <= value <= 100.
func MarginRightPercent(perc int) Option { _ = "STUB: not implemented"; return *new(Option) }

// MarginBottomPercent sets reserved space outside of the container at its bottom.
// The provided number is a relative margin defined as percentage of the container's height.
// Only one of MarginBottom or MarginBottomPercent can be specified.
// The value must be in range 0 <= value <= 100.
func MarginBottomPercent(perc int) Option { _ = "STUB: not implemented"; return *new(Option) }

// MarginLeftPercent sets reserved space outside of the container at its left.
// The provided number is a relative margin defined as percentage of the container's height.
// Only one of MarginLeft or MarginLeftPercent can be specified.
// The value must be in range 0 <= value <= 100.
func MarginLeftPercent(perc int) Option { _ = "STUB: not implemented"; return *new(Option) }

// PaddingTop sets reserved space between container and the top side of its widget.
// The widget's area size is decreased to accommodate the padding.
// The provided number is the absolute padding in cells and must be zero or a
// positive integer. Only one of PaddingTop or PaddingTopPercent can be specified.
func PaddingTop(cells int) Option { _ = "STUB: not implemented"; return *new(Option) }

// PaddingRight sets reserved space between container and the right side of its widget.
// The widget's area size is decreased to accommodate the padding.
// The provided number is the absolute padding in cells and must be zero or a
// positive integer. Only one of PaddingRight or PaddingRightPercent can be specified.
func PaddingRight(cells int) Option { _ = "STUB: not implemented"; return *new(Option) }

// PaddingBottom sets reserved space between container and the bottom side of its widget.
// The widget's area size is decreased to accommodate the padding.
// The provided number is the absolute padding in cells and must be zero or a
// positive integer. Only one of PaddingBottom or PaddingBottomPercent can be specified.
func PaddingBottom(cells int) Option { _ = "STUB: not implemented"; return *new(Option) }

// PaddingLeft sets reserved space between container and the left side of its widget.
// The widget's area size is decreased to accommodate the padding.
// The provided number is the absolute padding in cells and must be zero or a
// positive integer. Only one of PaddingLeft or PaddingLeftPercent can be specified.
func PaddingLeft(cells int) Option { _ = "STUB: not implemented"; return *new(Option) }

// PaddingTopPercent sets reserved space between container and the top side of
// its widget. The widget's area size is decreased to accommodate the padding.
// The provided number is a relative padding defined as percentage of the
// container's height. The value must be in range 0 <= value <= 100.
// Only one of PaddingTop or PaddingTopPercent can be specified.
func PaddingTopPercent(perc int) Option { _ = "STUB: not implemented"; return *new(Option) }

// PaddingRightPercent sets reserved space between container and the right side of
// its widget. The widget's area size is decreased to accommodate the padding.
// The provided number is a relative padding defined as percentage of the
// container's width. The value must be in range 0 <= value <= 100.
// Only one of PaddingRight or PaddingRightPercent can be specified.
func PaddingRightPercent(perc int) Option { _ = "STUB: not implemented"; return *new(Option) }

// PaddingBottomPercent sets reserved space between container and the bottom side of
// its widget. The widget's area size is decreased to accommodate the padding.
// The provided number is a relative padding defined as percentage of the
// container's height. The value must be in range 0 <= value <= 100.
// Only one of PaddingBottom or PaddingBottomPercent can be specified.
func PaddingBottomPercent(perc int) Option { _ = "STUB: not implemented"; return *new(Option) }

// PaddingLeftPercent sets reserved space between container and the left side of
// its widget. The widget's area size is decreased to accommodate the padding.
// The provided number is a relative padding defined as percentage of the
// container's width. The value must be in range 0 <= value <= 100.
// Only one of PaddingLeft or PaddingLeftPercent can be specified.
func PaddingLeftPercent(perc int) Option { _ = "STUB: not implemented"; return *new(Option) }

// AlignHorizontal sets the horizontal alignment for the widget placed in the
// container. Has no effect if the container contains no widget.
// Defaults to alignment in the center.
func AlignHorizontal(h align.Horizontal) Option { _ = "STUB: not implemented"; return *new(Option) }

// AlignVertical sets the vertical alignment for the widget placed in the container.
// Has no effect if the container contains no widget.
// Defaults to alignment in the middle.
func AlignVertical(v align.Vertical) Option { _ = "STUB: not implemented"; return *new(Option) }

// Border configures the container to have a border of the specified style.
func Border(ls linestyle.LineStyle) Option { _ = "STUB: not implemented"; return *new(Option) }

// BorderTitle sets a text title within the border.
func BorderTitle(title string) Option { _ = "STUB: not implemented"; return *new(Option) }

// BorderTitleAlignLeft aligns the border title on the left.
func BorderTitleAlignLeft() Option { _ = "STUB: not implemented"; return *new(Option) }

// BorderTitleAlignCenter aligns the border title in the center.
func BorderTitleAlignCenter() Option { _ = "STUB: not implemented"; return *new(Option) }

// BorderTitleAlignRight aligns the border title on the right.
func BorderTitleAlignRight() Option { _ = "STUB: not implemented"; return *new(Option) }

// BorderColor sets the color of the border around the container.
// This option is inherited to sub containers created by container splits.
func BorderColor(color cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// FocusedColor sets the color of the border around the container when it has
// keyboard focus.
// This option is inherited to sub containers created by container splits.
func FocusedColor(color cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// TitleColor sets the color of the title around the container.
// This option is inherited to sub containers created by container splits.
func TitleColor(color cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// TitleFocusedColor sets the color of the container title when it has
// keyboard focus.
// This option is inherited to sub containers created by container splits.
func TitleFocusedColor(color cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// splitType identifies how a container is split.
type splitType int

// String implements fmt.Stringer()
func (st splitType) String() string { _ = "STUB: not implemented"; return "" }

// splitTypeNames maps splitType values to human readable names.
var splitTypeNames = map[splitType]string{
	splitTypeVertical:   "splitTypeVertical",
	splitTypeHorizontal: "splitTypeHorizontal",
}

const (
	splitTypeVertical splitType = iota
	splitTypeHorizontal
)

// LeftOption is used to provide options to the left sub container after a
// vertical split of the parent.
type LeftOption interface {
	// lOpts returns the options.
	lOpts() []Option
}

// leftOption implements LeftOption.
type leftOption func() []Option

// lOpts implements LeftOption.lOpts.
func (lo leftOption) lOpts() []Option { _ = "STUB: not implemented"; return nil }

// Left applies options to the left sub container after a vertical split of the parent.
func Left(opts ...Option) LeftOption { _ = "STUB: not implemented"; return *new(LeftOption) }

// RightOption is used to provide options to the right sub container after a
// vertical split of the parent.
type RightOption interface {
	// rOpts returns the options.
	rOpts() []Option
}

// rightOption implements RightOption.
type rightOption func() []Option

// rOpts implements RightOption.rOpts.
func (lo rightOption) rOpts() []Option { _ = "STUB: not implemented"; return nil }

// Right applies options to the right sub container after a vertical split of the parent.
func Right(opts ...Option) RightOption { _ = "STUB: not implemented"; return *new(RightOption) }

// TopOption is used to provide options to the top sub container after a
// horizontal split of the parent.
type TopOption interface {
	// tOpts returns the options.
	tOpts() []Option
}

// topOption implements TopOption.
type topOption func() []Option

// tOpts implements TopOption.tOpts.
func (lo topOption) tOpts() []Option { _ = "STUB: not implemented"; return nil }

// Top applies options to the top sub container after a horizontal split of the parent.
func Top(opts ...Option) TopOption { _ = "STUB: not implemented"; return *new(TopOption) }

// BottomOption is used to provide options to the bottom sub container after a
// horizontal split of the parent.
type BottomOption interface {
	// bOpts returns the options.
	bOpts() []Option
}

// bottomOption implements BottomOption.
type bottomOption func() []Option

// bOpts implements BottomOption.bOpts.
func (lo bottomOption) bOpts() []Option { _ = "STUB: not implemented"; return nil }

// Bottom applies options to the bottom sub container after a horizontal split of the parent.
func Bottom(opts ...Option) BottomOption { _ = "STUB: not implemented"; return *new(BottomOption) }

// KeyFocusNext configures a key that moves the keyboard focus to the next
// container when pressed.
//
// Containers are organized in a binary tree, when the focus moves to the next
// container, it targets the next leaf container in a DFS (Depth-first search) traversal.
// Non-leaf containers are skipped. If the currently focused container is the
// last container, the focus moves back to the first container.
//
// This option is global and applies to all created containers.
// If neither of (KeyFocusNext, KeyFocusPrevious) is specified, the keyboard
// focus can only be changed by using the mouse.
func KeyFocusNext(key keyboard.Key) Option { _ = "STUB: not implemented"; return *new(Option) }

// KeyFocusPrevious configures a key that moves the keyboard focus to the
// previous container when pressed.
//
// Containers are organized in a binary tree, when the focus moves to the previous
// container, it targets the previous leaf container in a DFS (Depth-first search) traversal.
// Non-leaf containers are skipped. If the currently focused container is the
// first container, the focus moves back to the last container.
//
// This option is global and applies to all created containers.
// If neither of (KeyFocusNext, KeyFocusPrevious) is specified, the keyboard
// focus can only be changed by using the mouse.
func KeyFocusPrevious(key keyboard.Key) Option { _ = "STUB: not implemented"; return *new(Option) }

// KeyFocusSkip indicates that this container should never receive the keyboard
// focus when KeyFocusNext or KeyFocusPrevious is pressed.
//
// A container configured like this would still receive the keyboard focus when
// directly clicked on with a mouse or when via KeysFocusGroupNext or
// KeysFocusGroupPrevious.
func KeyFocusSkip() Option { _ = "STUB: not implemented"; return *new(Option) }

// FocusGroup represents a group of containers that can have the keyboard focus
// moved between them sharing the same keyboard key.
type FocusGroup int

// KeyFocusGroups assigns this container to focus groups with the specified
// numbers.
//
// See either of (KeysFocusGroupNext, KeysFocusGroupPrevious) for a description
// of focus groups.
//
// If both the pressed key and the currently focused container are configured
// to be in multiple matching focus groups, focus will follow the first
// focus group defined on the container, i.e. the order of the supplied groups
// matters.
//
// If not specified, the container doesn't belong to any focus groups.
// If called with zero groups, the container will be removed from all focus
// groups.
func KeyFocusGroups(groups ...FocusGroup) Option { _ = "STUB: not implemented"; return *new(Option) }

// KeyFocusGroupsNext configures a key that moves the keyboard focus to the
// next container within the specified focus groups.
//
// Containers are assigned to focus groups using the KeyFocusGroup option.
// The group parameter indicates which groups is the key attached to. This
// option can be specified multiple times to define multiple keys for the same
// focus groups.
//
// A key configured using KeyFocusGroupsNext only moves focus if the container
// that is currently focused is part of the same focus group as one of the
// group specified in this option. The keyboard focus only gets moved to the
// next container in the same focus group, other containers are ignored.
//
// The order in which the containers in the group are visited is the same as
// with the KeyFocusNext option.
//
// This option is global and applies to all created containers.
// Pressing either of (KeyFocusNext, KeyFocusPrevious) still moves the focus to
// any container regardless of its focus group.
func KeyFocusGroupsNext(key keyboard.Key, groups ...FocusGroup) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// KeyFocusGroupsPrevious configures a key that moves the keyboard focus to the
// previous container within the specified focus groups.
//
// Containers are assigned to focus groups using the KeyFocusGroup option.
// The group parameter indicates which groups is the key attached to. This
// option can be specified multiple times to define multiple keys for the same
// focus groups.
//
// A key configured using KeyFocusGroupsPrevious only moves focus if the
// container that is currently focused is part of the same focus group as one
// of the group specified in this option. The keyboard focus only gets moved to
// the previous container in the same focus group, other containers are
// ignored.
//
// The order in which the containers in the group are visited is the same as
// with the KeyFocusPrevious option.
//
// This option is global and applies to all created containers.
// Pressing either of (KeyFocusNext, KeyFocusPrevious) still moves the focus to
// any container regardless of its focus group.
func KeyFocusGroupsPrevious(key keyboard.Key, groups ...FocusGroup) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Focused moves the keyboard focus to this container.
// If not specified, termdash will start with the root container focused.
// If specified on multiple containers, the last container with this option
// will be focused.
func Focused() Option { _ = "STUB: not implemented"; return *new(Option) }
