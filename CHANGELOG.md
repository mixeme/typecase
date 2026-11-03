# Changelog

All notable changes to Typecase are documented here. Versioning follows
[semantic versioning](https://semver.org/). This file tells consumers what to
do, not what the authors did.

## Unreleased

### Dependencies

- What a consumer imports needs two modules outside the standard library,
  and only these: `golang.org/x/text` (`widget.Slug`) and
  `golang.org/x/net` (`page.Check`); both are direct requirements now.

### Breaking

- `Push` is no longer a `SheetFootPart`: a sheet's Cancel stands in a
  pushed `Line` with its `Kbd` «Esc», as in every sheet (Р-81).
- A public page's `Banner.End` takes a `Line` only, with one button too
  (Carrel's task 234: one rule for the banner's end).
- **`Runbar.Every` is a `time.Duration`**, printed in milliseconds
  (`data-tc-every`): a poll every 1500 ms can be said (Carrel's task 226).
  The old `Every: 5` still compiles and means 5 ns — it prints
  `data-tc-every="0"`, and the page polls at its default 5000 ms; write
  `Every: 5 * time.Second`.
- **`DocTasks.Title` is a string** drawn as the document's heading
  (`.tc-doc__heading`), not a `Rubric`: a note's Markdown heading is the
  author's text, not the interface's rubric (Carrel's task 196).
- `MainFoot.Roomy` is gone with its rule (`.tc-main__foot--16`): every
  public page ends with the same foot, the link's day only (Carrel's task
  196).

- `PhotoCrop.After` is gone: the crop page says nothing after its fields
  (Carrel's task 197, Р-77).
- **A `Mark` is no `ValuePart`**: the collection a value came from is a
  `Dim` with its `Mark` (`Dim.Mark`, new) — the mark and the name stay on
  one line (`.tc-dim:has(> .tc-mark--inline)`; Carrel's task 190).
- **«Save as» names its element by `data-tc-save-from`**, not
  `data-tc-from` (`Button.SaveFrom`; Carrel's task 224): `data-tc-from` is
  a derived field's source, a field's name (`Input.From`). The page check
  looks that name up in the field's form, as it does `data-tc-disables`.
- `TimeRow.Short` is gone with its rule (`.tc-row--timecols--time`): a
  search's event shows its day in the lead, the time beside the place
  (Carrel's task 189). `Say.Mono` is gone with `.tc-say--mono`, and a
  `Runbar` no longer takes a `Line`, nor a `Line` a `Say`: a search's bar
  has no «answered · Search again» of its own — the header's button is the
  one door, as on every list.
- `Chevron` (the small arrow after a cell's value) is gone with its rule
  (`.tc-ico--inline`): a cell's value that can be changed is a closed
  choice (Carrel's task 212).
- `MainTree` is gone with its rule (`.tc-main__tree`): davloom's tree is
  a sheet now, behind «What the client sees» (the owner, 19.10).
- `RubricGap13` is gone with its rule (`.tc-rubric--m13-0-6`): «What it
  opens» is the first rubric of its card now (Carrel's task 212).
- `FieldRow`, `MarkH18` and `Mark.Focus` are gone, with their rules
  (`.tc-field__row`, `.tc-mark--h18`, `.tc-mark.is-focus`): the label
  colour of a feed is a choice now, `Swatches` with `Name` (Carrel's task
  211), and no frame draws marks in a field any more. `Mark` is no part
  of a `Field`.
- `Turn` (the arrow at a card's foot) and `Card` as a part of a
  `FieldSet`, and `Badge` as a part of a `Check` are gone: no frame draws
  them since the backup job lost its trigger card (Carrel's task 193).
- `MarkH19` is gone: no frame draws a mark of 19 px since a record's head
  carries its collection's mark beside the collection's name.
- **A field of a date or a time is as wide as its value** (the owner,
  16.10): not narrower and not wider — its value and its icon, whatever
  the size class of the field around it; the field around it takes that
  width too. A date of 140 px with room to spare is now about 125 px; on
  a phone a date field no longer stretches across the column. The width
  ladder asks it both ways (a new check «wide»).
- **A check of one of several (`Check{Radio: true}`) is a circle**, the
  chosen one with a dot: a square promised that the next one could be
  ticked too (the owner, 16.10). 56 of them in the mockup changed so.
- **A widget holds its own; a container does not take its minimum** (the
  owner's decision of 15.10, Р-65). On a phone (a column of 460 px or
  less) a bar, a document's bar, a range inside itself, the buttons of a
  table's row in their cell and a sheet's foot wrap what does not fit after
  the ranks and «⋯», instead of squeezing it or cutting it at the edge. A
  bar is one row while it fits. A field of a date or a time is never
  narrower than its value and its calendar (`min-width:max-content`); a
  filter field (a field with a search line inside) keeps its 88 px itself.
  A search of the top bar with an open query takes a row of its own on a
  phone, its query seen — its rules never applied before (a weight that
  lost to the folded search). On a phone a small button is 38 px like the
  rest. The cell of a row's buttons stays a table cell: the row's line no
  longer breaks at it.
- **The live layer sends a form that is not GET once** (Carrel's note 58):
  when its answer has come, it is not sent again the plain way — an answer
  sent on to another address goes there, any other answer that is no part
  (an error, a page without the zone) is shown whole, as the plain way
  shows it. Only a request that had no answer goes the plain way. A link and
  a GET form still go the plain way when the answer is no part.
- `Pick` prints `value` on every option, the empty one too: an option
  without it sent its words — «Do not rewrite» went as those words, not as
  the empty value the server waits for.
- `TimeRow.Caption *Badge` is now `TimeRow.Badges []Badge` (the owner's
  decision of 14.10, Р-61): pass all the event's badges by importance —
  `AgendaRow.Badges` as it is; the row shows those that fit whole, two side
  by side («4 attendees» «repeats weekly»), a badge that does not fit is
  not drawn rather than cut, and the cell's hint names them all when there
  is more than one. The badges column of `tc-row--timecols` is 168 px at
  every width where it stands (it was 96 and 88).
- `Thumb` with `Src` prints its icon too, under the picture (Carrel's task
  148): pass `Icon` with `Src`, as for a file without a thumbnail — a `Thumb`
  that has only `Src` prints an icon with no name.
- `DueRow.Tags` and `ExcerptRow.Tags` are gone: a row of a list has no tags
  by default in any kind (the owner's decision of 10.10, Р-53) — they stand
  in the record's panel and in the rail. Tags in a row are the column a
  person turns on: `ColumnRow` with `ColumnTags`, whose cell now shows the
  first of `ColumnCell.Tags` and counts the rest («работа +1»), naming them
  all in its hint. Pass all the record's tags; put the one the list is
  filtered by first. The grids of `tc-row--tickmarkcols` and
  `tc-row--excerpt` have no tags column.
- `RailFoot` has no state «about» any more: no frame draws a rail foot of
  the build line alone since the duplicates frames got the davloom mark
  (Carrel's task 136). The widget prints what it printed.
- `Badge.Icon` is gone: no frame draws a badge with an icon since the
  «linked» badge left the tags column of the notes list (Carrel's task 105).
- `TickRow.Trail` no longer takes a `Button`: a refused source has no Retry
  in the rail any more (the owner's decision of 02.10, below).
- `Sec` no longer takes a `Table`: no frame puts one in a section since the
  attendees became a `SecList` with `Columns` (Carrel's task 77).

### Added

- `SecHint` is a `DocmetaPart` too: a section of a document's side column
  says under its rubric what it keeps (Carrel's task 253: «Kept as the
  server sent them» of a note).
- `frame2go -self -out DIR` writes each frame's draft to
  `DIR/<frame>/draft.go`, a package per frame (Carrel's task 245): many
  frames are one call — the mockup and the library are read once, all 294
  route frames take about six seconds.
- `Button.Busy` — what a submit button says while its form is being sent
  (`data-tc-busy`): with the library's script it takes these words and is
  off until the answer, its icon kept, its name and value still sent; a
  page come back from the history has it as it was (Carrel's task 242).
- **The reading column of a settings page** (the owner, 27.10):
  `FieldW580` — a field as wide as the column; `HintColumn` — a field's
  hint at most 580 px (`.tc-field__hint--max580`); `Msg.Column` — a message
  at most 580 px (`.tc-msg--max580`). Tables stay full width.
- `Shell.Dialogs` — the page's own sheets without an address, each closed
  under its ID until a button (`Button.Opens`) or a link to «#ID» opens it
  with the library's script: the keys sheet of «Keyboard shortcuts» and
  «?» (Carrel's task 239). The frames draw them in the zones' scheme
  (`shell-1`) only, as the place of a sheet.
- `MenuItem.Required` — a radio item of a choice the form cannot be sent
  without, without a script too: the notebook of a quick note not chosen
  yet (Carrel's task 238).
- **`shape` reads the zone of messages as its messages** (Carrel's task
  238): a screen whose messages stand in `tc-notices` has the shape of a
  frame that draws them in the page; an empty zone is nothing.
- **Paper for every kind** (Carrel's task 228): `PrintRow.Field` — a
  card's field, its name then its value; `PrintRow.Head` — the row naming a
  table's columns, small capitals over the rows (the owner, 25.10);
  `PrintText` — a note's paragraph or heading in reading width;
  `PrintTick` (`PrintCell.Tick`, `PrintText.Tick`) — a drawn box, ticked
  when done, the done text struck through the middle of its small letters.
  Three-column rows keep one width per column down the sheet and fold
  under the name on a narrow sheet. A `Top` without `Actions` prints no
  actions block (the offline page).
- `Filter.Clear` — the same address without the filter: while the line
  holds a word, a cross at its end leads there (`.tc-in__clear`), without
  a script too (the owner, 25.10; Carrel's task 231).
- **The live layer for Carrel's zones** (Carrel's task 226):
  - an answer carries the page's forms after its zone's widget
    (`page.Fragment.Forms`, `ZonePart.Forms`): the widget takes the zone's
    place whole, each `tc-post` form takes the place of the page's form of
    the same ID or joins the end of the page;
  - `Runbar.Closed` — an empty, hidden line that waits for a run; a hidden
    line is not followed, so a line emptied by the event that ends its run
    no longer opens its stream again;
  - `Notices` — the zone of a page's messages in `Main`, without a look of
    its own (`.tc-notices{display:contents}`), hidden when empty; it holds
    messages and a message's technical block (`NoticesPart`: `Msg`,
    `MainDiag`; Carrel's task 235);
  - `Shell.Sheet` — the place of a sheet over the whole shell (the top
    bar, the rail and the page dimmed, as `Overlay`), closed until an
    answer fills it.
- `DocRef.From`, `DocAttach.From` — the copy a line of a document's side
  column comes from, on a linked record (Р-78; Carrel's task 230).
- `MenuText` — a menu row that only says something, a name over a line
  (`.tc-menu__text`): how to mount a public folder on each system. A
  public page's address `Strip` takes a `Drop` (Carrel's task 196).
- **The search line's suggestions come from the server** (Carrel's task
  227): `page.Document.Suggest` names their address (`data-tc-search-suggest`
  on the root); with the live layer the line asks it `?q=…` as it gets the
  focus and at each pause in typing, and puts the answer — a `Menu` whose
  items say `When` — under the line, where the core filters it, walks it
  and hides it as a printed one. The page prints no menu; without the live
  layer there is none. `TestSearchSuggestAsks`.
- **`Field.Post` and `Field.Sends`** — a field that is a form of its own
  (`Field.Action`) can be a POST with hidden fields: the duplicates'
  threshold, kept with the person's decisions (Carrel's task 190).
  **`FaceRow.SourceHint`** — a linked person's row names its sources in
  full, as a linked task's and note's do. **`Dim.Mark`** — see Breaking. The
  close button of a panel whose head is centred stands at the top, level
  with the name (`.tc-head--mid > .tc-acts:has(> .tc-close)`; `s-contacts-1`,
  `-7`, `s-notes-1`). **A row's title with a badge** prints its words in
  `.tc-row__words`: on a narrow list they shorten with «…» and the badge
  stays whole (twenty titles in the mockup).
- **`SizeRow.Mark`** — the collection's mark in the tick's place, where
  nothing is done with the file but open it: a search's results (Carrel's
  task 189). **`Pad` takes a `Menu`**: the search page's recent searches.
- **A button after a message's words stands apart from them**
  (`.tc-msg > .tc-btn`, `.tc-msg__aside > .tc-btn`: 6 px to the left), and
  a link button in an alert keeps the button's ink instead of the alert's
  red (`.tc-msg.is-alert > a.tc-btn`) — the owner, 22.10 (Carrel's task
  222). Ten frames gain the gap; `s-conflict-10`, `s-io-19`,
  `s-profile-32` lose the red.
- **`TagAdd.Enter`** (`TagEnter{Name, Value, Label, Of}`) — the unseen
  default button of a tag field's line (`.tc-in__enter`): without a script
  Enter in the line sends the form with it — the letter's «add the
  recipient», `op=add` — not with the form's first submit (Carrel's task
  221). With the library's script Enter makes a tag and sends nothing.
- **`Query.Of`** and **`PostForm.Method`** — a picker's search in a sheet that
  is a form of its own goes with a GET form of its own (Carrel's task 221).

- **`Dialog.Ranges`, `Dialog.RangesOf`** — a step sends parts of the file an
  earlier step kept in the page (Carrel's task 220, Р-71): the live layer
  keeps a `FileField.Ranges` file under its field's name; a dialog naming
  that field asks the ranges address with its own fields — the chosen
  collections — and sends the parts named then. A form that names a zone
  (`Main.Zone`, `Main.Target`) takes the answer of a ranges send in place,
  so the page and its file stay. Without the layer the whole file goes.
- The frame generator reads a page column's `Zone` and `Target`.

- **The answers about taken names go before the files** (Carrel's task
  219): `taken.<name>` stands before the file field — with its `form`, as
  the sizes do — so a server that reads the body as a stream has the answer
  when the file comes. An earlier asking's answers are dropped.
- The frame generator reads a queue's `Progress`, `Zone`, `Row` and `Taken`.

- **`PickOption.Disables`** — an option of a pick switches fields of its
  form off while it is chosen (`data-tc-disables` on an option; Carrel's
  task 218): «None — up to 16 MB» and the folder for a large archive.
- **`Doc.Multipart`** — an edit's document sends the files that fields
  outside it pick for it (`FileField.Form`): Attach in a note's edit bar
  and column (Carrel's task 249). A file field inside makes it so by
  itself.
- **`MenuItem.To` with `Form`** sends that form from outside it, as
  `Button.To` with `Form` does: «From Files» in a new note's bar
  (Carrel's task 247).
- **`PostForm.Multipart`** — a post form sends multipart without a file
  field of its own: the answer to a taken name (Carrel's task 218).
- `Dialog` with a file field among its parts sends multipart, as `Form`,
  `Sheet` and `Card` do (Carrel's task 218).

- **`SegOption.Disables`** — an option of a switch with a `Name` switches
  fields of its form off while it is chosen (`data-tc-disables` on a radio;
  Carrel's task 194): «The whole calendar» and the days of a range. The
  core script goes over the whole group when another option is chosen.
- **`Form.Recount` and `Dim.ID`** — a form asked again as it changes: with
  the live layer a field that changes asks the form's address with its
  fields (GET) and puts the elements it names in place — the line at Attach
  that says what the chosen days weigh (Carrel's task 194). A field that is
  off is not asked with. Without the layer the line says so after Attach.
- **`SheetWidth900`** — a picker over a screen: a rail and a list (the
  letter's attach sheets).
- **`DocFile.Href` drafted from a frame** — a file published with its note
  is a link on the public page.
- `Sheet` with a form among its parts is printed plain: a form holds no
  form (the calendar days to attach).

- **`SegOption.Form`, `SegOption.To`** — an option of a switch that sends
  a form to its address, without a script too (Carrel's task 215): a
  davloom device's access, read only or read and write, is a POST each.
- **`FileField.Taken` and `Queue.Taken`** — the taken names asked before
  the files go (Carrel's task 214: the server keeps nothing of a file while
  a question waits). With the live layer a form whose queue shows it asks
  the field's address first — its hidden fields and a «name» a file — and
  the server answers the taken names, a line each; the queue asks about
  each with a row of its `Taken` template (buttons valued keep, replace,
  skip); when all are answered the skipped files leave the field, the
  answers go as «taken.<name>», and the form goes with its progress.
- **`FileField.Ranges`** — a file that goes in parts (Carrel's task 213 p.
  2: a large archive is read where it lies, never copied whole). With the
  live layer a form with such a field asks that address first — its other
  fields, «<name>.name» and «<name>.size» — and the server answers lines
  «offset length»; the form then goes to its action with those parts of
  the file under the field's name, each called «offset-length», and the
  answer is the page. Without the layer, or when the asking fails, the
  whole file goes. The restore step `s-backup-15`, `-16` uses it
  (`/app/settings/backup/restore?op=ranges`).
- **`Sheet.ID` on a sheet that is no form** (Carrel's task 210): the id
  a link's swap names — the file picker's servers.
- **`SlimRow.Reason`** — the tick of a row off, and why, the aside as an
  alarm: an account whose saved password the server refused cannot go
  into a device's profile (Carrel's task 211).
- **`Check.Reason`** — a check of a form that is off, and why, as a
  button's reason: the input is disabled, the reason the label's title,
  the check dimmed (`.tc-check:has(> .tc-check__box:disabled)`). «Let
  users send letters» while mail is not configured (Carrel's task 208).
- **`widget.BarForm`** — a bar's filter that is a GET form (Carrel's task
  195): its fields (a pick with a menu of radios, a date, a dim word) stand
  in the bar as its own; with the library's script a change sends it,
  without one its Show does (shown only then). A bar holding one wraps
  what does not fit at any width. A menu's radio item always carries its
  value, an empty one too («All actions»).
- **`Input.Of`** — the form a field belongs to without standing in it, as
  `MenuItem.Of` and `Send.Of`: an address in a head beside the button
  that sends it.
- **`FieldGridColumn`** — a column of a grid of choices: a source and
  what is chosen under it, together however the grid wraps.
- **`CardFields.Top`** — a card's row of fields aligned by their tops:
  refusals under some do not lift the others; a button stands in a
  `Field` with an empty label.
- `MsgFlushY8` — a message 8 px from what stands above and below it,
  none at the sides: a refusal between a field and a tree in a card.
- **A sheet's head takes a `Drop`** (`Drop.Place` pushes it to the end),
  and **`MenuItem.Swap`** puts a link item's answer in place with the
  live layer, as a tree's node does; a tree's link row carries its `Swap`
  too — the connection of a sheet's tree chosen in its head (Carrel's
  task 205).
- **`Check.Of`** — the same for a check: the consent to the DAV
  validator's second stage under its trace.
- **A button that waits for its form — `ButtonWaits`, `Check.Required`**
  (Carrel's task 198). A radio with `Required` must be chosen before its
  form goes; a button of `ButtonWaits` looks disabled while a required
  field of the form, dialog or sheet around it is empty, and the browser
  refuses it then, naming the field; its `Reason` is its title. What to
  choose the screen says above the fields, not beside the button (the
  owner, 17.10). All by CSS (`:has(:invalid)`): it works without a script.
- A `Msg` takes a `Mark` (a collection named in a message); a `Dialog`
  takes a `Diag` (what was asked and answered, under a message).
- **`MenuItem.Of`** — the form a radio item of a menu belongs to without
  standing in it (Carrel's task 203): the collection chosen in an edit's
  head stands outside the edit's form, and without `form="…"` the browser
  sent none of it. A menu item with `Name` outside a form that sends is
  right when it names its form with `Of` (`docs/guide/screen.md`, «Проверка»);
  `page.Check` asks the ID to name a form on the page.
- **`widget.Slug`, `Input.From`, `SlugOf`** — a field derived from another
  (Carrel's task 200): a new collection's address from its name. `Slug` is
  the one rule (NFD, the Cyrillic table, lower case, spaces and hyphens to
  one hyphen, `[a-z0-9-]`, empty — `col`) for the server; the library's
  script derives the same as the name is typed, into the field's
  placeholder and the `SlugOf` pieces of its hint; what is typed in the
  field wins. Without a script the field shows the placeholder the server
  printed. Typecase now requires `golang.org/x/text` (NFD).
- **`page.Document.Theme`** — the theme the server chose for this browser
  (Carrel's task 186; a cookie): «light» or «dark» is printed on `<html>`
  and shows without scripts, «auto» leaves it to the system; theme.js and
  the theme's switch keep the server's choice and neither read nor write
  what the browser remembers. Empty, the library decides as before.
- **A `Field` in an edit's bar (`Docbar`)**: where a note is saved, after
  Cancel (the owner, 16.10).
- **`Field.Changed`** — a field that a change brings: printed hidden, it
  shows while one of the named fields of its form differs from what the
  page brought and goes when the edit is undone; a field the page brought
  shown is left alone. Without the script the screen asks after Save
  (the owner, 16.10: the question about published links comes at once).
- **`SecFile.Terms`, `SecFile.Change`** — a file of a letter that goes as
  a link: the link's terms in words and the button that opens the publish
  sheet over the letter (Carrel's task 165).
- **`Seg.Tertiary`** — a switch of the third rank: it goes on a phone where
  it means nothing (a note's width of text, narrower than the reading
  width anyway).
- Icons `i-import` and `i-export` (a page with an arrow into it and out of
  it), `i-upload` (an arrow up out of a tray, the twin of `i-download`),
  `i-code` (angle brackets: a record's source).
- **Checks of what a widget promises, not of the frame** (Р-65). The width
  ladder now asks every field of a date or a time, button, option of a
  switch, tag, badge and pick whether it holds its content, and every
  control whether an ancestor that cuts hides a part of it — measured on
  the widget, so a frame drawn wrong is red too. The gate `context-minimum`
  refuses a rule of a context that sets another control's min-width to
  zero or cuts it; the one exception (the folded search) is listed with
  its reason.
- **`Say{Full, Short, Mono}`** — words in two lengths: the short ones show
  where the column is 640 px or less, to say it shorter instead of
  squeezing it («Done · 13 of 14 sources answered» → «13 of 14»).
- **`TableCol.Main`** — the column that matters (Р-66): on a phone a table
  with one is cards — the name and «⋯», under them this column named by
  its label; the other columns go, the row's buttons are in «⋯».
- `Span.Title` and a `Span` in a table's cell: a failure in short words,
  said in full in its hint («1 not copied»).
- **A page checks how its parts name each other** (the owner's thought
  after the review: the constructor must not break when it is put
  together). `page.Check(html)` and `Document.Check()` read a whole printed
  page as the browser and the library's script will and return a `Problem`
  for every name that does not hold: an id given twice; `form`, `list`,
  `for`, `aria-*`, `data-tc-progress`, `data-tc-add`, `data-tc-open`,
  `data-tc-swap`, `data-tc-recount` and the other ids a part gives — naming
  nothing, or the wrong thing (`data-tc-drop` and `data-tc-paste` want a
  form with a file field, `data-tc-add` a template, `for` a control);
  `data-tc-disables`, `data-tc-follows`, a field's `data-tc-changed`
  (a stream line's is the name of its event), `data-tc-day`,
  `data-tc-slug-of` and an option's `data-tc-shows` — naming no field of
  their form; `data-tc-target` and the document's address zone —
  naming no zone on the page; a file field whose form, around it or named,
  does not send multipart — a browser sends the file's name only; a
  submit button of no form (a sheet whose
  body holds a form of its own prints no form of its Action: `Action`,
  `Method`, `Fields` and `Target` of the sheet are dropped, and the button
  does nothing). Names with «{n}» are the template's and are left. Such
  a page renders and looks right; what breaks is what is pressed. Run it in
  your tests over the pages you print. `TestCheck`.
- **A queue folds in place** (a tail of Carrel's task 88): `Queue.Fold`
  (`QueueFold{Hide, Show, Total}`) — with the library's script its Toggle
  hides the list and says Show, the total says `Total` («{percent}%») once
  the live layer knows the upload, and pressed again it unfolds, with no
  request; without a script Toggle is the server's link. The live layer
  gives a queue's title and total «{percent}» too. The mockup's folding
  queues carry it.
- **A button counts what its form's choices write** (the owner's decision
  of 14.10, Р-62): `Button.Count` («Rewrite {n} records») and
  `Button.CountOne` («Rewrite 1 record»), with `Tally` in the sheet — one
  record a line, each the conditions on which it is written: `name=value`
  (a tick or choice of that name has that value), `name` (that field is not
  empty), `a|b` (either), all of them at once. `tc.js` says the count again
  at each change of the form and turns the button off at none; without a
  script the label stays as printed. In «Rewrite the addresses» (s-contacts-29)
  a record is the address it names and its groups — `to.1 group=scheduling
  group=others` — so a record of two groups counts only while both are
  ticked, as the sheet says.
- **`shape` does not count an element with `hidden` as drawn** (Carrel's
  task 161): a screen may print a field hidden — the fields «Custom…» at a
  reminder of a ready step — that its frame does not draw, and a frame may
  draw one the screen does not print. Kept are the places a script opens
  or fills, which the frames draw hidden so that the screen has them: the
  progress bar (`tc-runbar`), a menu (`tc-menu`), the scrim. Shown, such a
  field is an extra element, as before.
- **A row added from a template numbers what its options show** (Carrel's
  task 161): `tc.js` replaces «{n}» in `data-tc-shows` of the row's options
  (`PickOption.Shows: "reminder.{n}.count …"`) as in the names of its
  fields, and sets the added row's fields by its choice — «Custom…» of an
  added reminder opens that reminder's fields, which are sent under its
  number; until then they are hidden and not sent.
- **`ColumnBadges`, `ColumnCell.Badges`**: the agenda in chosen columns
  keeps the badges (Р-61) — the column «Attendees and status» is
  `ColumnBadges` with the event's badges, as `TimeRow`'s, 168 px, gone at
  620 px; an event with nothing to say has an empty cell. `ColumnStatus`
  stays a task's status in words.
- **Zones: a part of the page changes without a page load** (Carrel's task
  84, the engine). A link or a form inside a widget that names a `Target`
  asks, with the live layer, the same address as without it, with the
  headers the document names — `page.Document.Zones{ViewHeader, ViewValue,
  TargetHeader, TokenField, TokenHeader, Address}` — and the answer is put
  into the zones: `page.Fragment{Body, Extra}` prints the part asked for
  and, after it, each other zone in a `<template data-tc-zone>`. A zone is
  a widget the application named: `Zone` on `Main`, `Side`, `Top`,
  `Runbar`, `Queue`, `Scrim`; `Target` on `Nav`, `List`, `Bar`, `Form`,
  `Sheet`, `Button`. A widget that carries the zone's name takes its place
  whole, anything else becomes its content; an empty answer empties and
  hides it. An answer that is not HTML, not 2xx or redirected sends the
  page to the address the plain way; a form's token goes into the header
  the document names; the zone named in `Zones.Address` puts a link's
  address into the history. A form that sends files still loads its answer
  whole, with its progress. Without the layer everything is a plain link or
  form. `TestZones` (browser and `page`).
- **A search's suggestions follow what is typed** (Carrel's task 92).
  `MenuItem.When` says when an item of the menu under the search line
  (`Top.Suggest`) is shown: `WhenEmpty` — only while the line is empty
  («Type to search everything»), `WhenTyped` — only while something is
  typed, «{q}» in its words being what is typed («Show all results for
  «{q}»»), `WhenFound` — only while its words hold what is typed, which is
  set bold (a jump to what the page already holds); an item without `When`
  — a recent search — always. Print the menu whole: every jump the page
  can offer. With the script the menu shows while the line or the menu has
  the focus, ↓ and ↑ walk it, Esc hides it; Enter sends the search's form.
  Without a script it stands as printed. `TestSearchSuggest`.
- **Installed as an application** (Carrel's task 91). `page.Document.App`
  names the application's manifest, the colour of the browser's bar
  (`theme-color`) and the icon — the head prints them; the manifest and
  the icons are the application's files. With `App.Offline`, the address
  of the page shown when the network is gone, tc.js registers the
  library's worker (`Assets.Worker`, `js/tc-sw.js`, served by `Static`
  with `Service-Worker-Allowed: /`; scope `App.Scope`, «/» by default). It
  keeps the stylesheet and the scripts the page loaded from the library's
  base, the fonts as they are asked for, and the offline page — nothing
  of the application's: pages and data always come from the network, and
  a page that cannot load is the offline page. `At` names the worker too.
  Serve the library at a new base for a new version: the page registers the
  worker anew, the new one takes over and drops the copies kept outside its
  base. With a manifest that has a name, `start_url`, `display` and an icon
  of 512 px, the browser finds nothing missing to install the page.
  `TestAppWorker`, `TestAppWorkerUpdates`, `TestAppInstalls`.
- **What needs a script, and the words when there is none** (Carrel's task
  90). The mockup's CSS has the rule opposite to `tc-nojs`: `tc-js` is not
  shown while the page runs no script. `Msg.ScriptOnly` and `Msg.NoScript`,
  `Button.ScriptOnly`, `MenuItem.ScriptOnly` print them. `s-files-1` shows
  both: the line about dropping onto the window with a script, and «Dropping
  files onto the window and the upload queue need scripts, and this page
  runs none. Upload sends the files; a name that is already taken is
  refused rather than overwritten.» without one (accepted by the owner on
  14.10). `TestScriptOnly`.
- **The theme's switch shows what this browser keeps** (Carrel's task 89,
  the theme). A `Seg` with `Theme` shows on load the choice kept under
  `tc-theme` — «Auto» when nothing is kept — whatever option the server
  printed as chosen: print any. `TestThemeShowsKept`.
- **A pasted file, and an upload that stays on the page** (Carrel's task
  88, without the drop on a folder). `Area.Attach` names the form whose
  file field takes a picture pasted into the text — it becomes a file
  picked there; words are pasted as words. Files pasted outside a field go
  where dropped ones would: to the form of the page's drop sheet
  (`Scrim.Drop`). A dropped or pasted file now sends its form by the
  form's submit event (`requestSubmit`), so the live layer can take it.
  `PostForm.Target` names the zone such a form asks for: with the live
  layer a form that sends files and has a target goes from the page — the
  same headers as any form of a zone — and its answer is put into the
  zones. `Queue.Progress` names the form whose upload the queue shows and
  `Queue.Row` is the row every file gets: the words «{name}», «{size}»,
  «{percent}» in the row and «{count}», «{done}», «{sent}», «{total}» in
  the title and the total are written as the upload goes, and the rows'
  meters move. A queue with a `Progress` and no `Items` is printed hidden.
  One request carries all the files, so a row's share is counted from the
  bytes sent so far. Without the layer the form loads its answer.
  `TestPasteAndQueue`.
- **A file dropped on the open card** (a tail of Carrel's task 88, frame
  `s-io-2`). `Side.Drop` names the form with a file field a file dropped
  on the panel is sent with (`data-tc-drop`). With the library's script a
  file over the panel marks it (`tc-side--drop`) and keeps the page's drop
  sheet away — the open card wins over the page; dropped there, the file
  goes to the panel's form, off the panel to the sheet's. Without a script
  nothing is dropped: the card's «Attach» stays. `TestSideDrop`.
- **Shortcuts are data** (Carrel's task 87). `Button.Key`, `MenuItem.Key`
  and `Query.Key` name the key that presses the button or the item, or puts
  the focus into the search line (`data-tc-key`: «n», «/», «?», «Ctrl+S»,
  «Ctrl+Enter»); the library has no keys of its own besides Esc, Tab and
  the arrows. A letter or a sign does nothing while the cursor is in a
  field or a menu is open; a chord with Ctrl (Cmd) works in a field too,
  for a button of that field's form only. Under an open sheet nothing is
  pressed. An item of a closed menu is pressed as well: «?» on «Keyboard
  shortcuts» opens the sheet its `Href` names. A button or an item that
  destroys (`ButtonDanger`, `ButtonDangerSolid`, `MenuToneDanger`) does not
  print its key. On a layout without Latin letters the key's place names
  the letter. In a `List` the arrows, Home and End move the focus along the
  rows that are links (or whose title is one); Enter opens the row, as a
  link does. Tab still visits every row. `TestShortcuts`.
- **A run's line follows its run** (Carrel's task 85). `Runbar` that is a
  zone names a stream of events — `Stream`, with the events' names
  `Changed` and `Ended` — and a poll address, `Poll`, asked every `Every`
  seconds. With the live layer the line takes each event's fragment (the
  line as it is now, and whatever other zones changed) and closes the
  stream on the event that ends the run; with no stream, or one that
  refused or broke, it asks the poll address for as long as the line that
  comes back names one and is shown. One run is followed once, whatever
  comes to replace its line. Without the layer the run shows after a
  reload. `TestRunFollows`: a stream, a poll, a stream that refuses, and
  the page without the layer.
- **Zones, the rest** (Carrel's tasks 84 and 83). The rail is a zone:
  `ShellBody.RailZoneName`, and `RailTarget` for its links. `Target` stands
  on the areas too — `Main`, `Side`, `Top` — and on `Menu`, `Seg`, `Tree`,
  `Table`, `Crumbs`: the nearest ancestor that names a target decides.
  `Side.Closed` prints a panel that waits for its answer; the shell's body
  makes room for a panel that arrives and takes it back when it is closed
  or emptied. **A zone the mockup does not draw at this width is not
  filled**: the side panel on a narrow screen — the link or form goes the
  plain way, to the page that shows the thing whole. `widget.PrintMain`
  prints parts of a page without the page, and `page.Printed` carries them
  in a fragment: a message with its technical block (`MainDiag`) for a zone
  of messages. Bars and heads that arrive later are watched for their width
  as those of the first page are. `TestZonePanel`.
- **A sheet that comes as an answer, and a Cancel that closes** (Carrel's
  task 84). `Scrim` with a `Zone` is the layer of sheets: a link or a form
  whose target is that zone gets the sheet as its answer, the sheet lies in
  the layer, the focus stays inside it, and it is closed by Esc, a click on
  the scrim or its Cancel without a request. `Button.Closes` makes a button
  close the sheet or panel it stands in (`data-tc-close`); without a script
  it is what it is besides — a link back or a plain button. Before, a
  sheet's Cancel had no behaviour of the library at all: only the «×»
  button, the scrim and Esc closed. `TestZoneSheet`.
- **A dialog that is a form** (Carrel's task 149). `Dialog.Action` (with
  `Method`, `Fields`) prints the dialog as a form: a step of a wizard whose
  fields stand in several cards sends them all. A form holds no form — a
  dialog one of whose parts is a form of its own is printed plain, and that
  part sends. `Send.One` makes a row's tick a radio button: one of several
  under the name (the archive to restore from).
- **An option of a pick that shows fields** (Carrel's task 156).
  `PickOption.Shows` names the fields of the form that are shown only while
  the option is chosen (`data-tc-shows`): «Custom…» of a reminder opens its
  number and unit. A field that is not shown is printed with `Field.Hidden`
  (`.tc-field[hidden]`); `tc.js` shows and hides them as the choice changes
  and switches a hidden field's controls off, so it is not sent. Without a
  script the fields stay as printed and the server answers with the form as
  chosen — it must not read the fields of an option that is not chosen.
  `TestPickShows`.
- **A tree's node that is a tick and has an address** (Carrel's task 158).
  `TreeRow` with `Send`, `Href` and `Swap`: the node stays a radio button of
  its form, `Href` is the address of the page drawn with this node chosen
  (`data-tc-href`), `Swap` the IDs of what to change (`data-tc-swap` — the
  sheet's `ID`). The live layer asks the address when the node is chosen
  and puts the sheet in place — the folder opens, the button that could not
  act can; the address does not go to the history. Without the layer the
  form's «Open the folder» button does the same through the server.
  `TestTreeNodeSwap`.
- **The selection bar sends without a script** (Carrel's task 148; the
  owner's decision of 09.10: without scripts the bar stands always). The
  ticks of a table's rows and of tiles are checkboxes of a form they do not
  stand in: `Send.Of` is the ID of a `PostForm`, printed as `form=` on the
  checkbox (`TableCell.Send`, `Tile.Send`). No form wraps the list — the
  bar under it holds forms of its own, and a form holds no form. The bar's
  buttons send that form with which action was pressed (`Button.Form`,
  `Name`, `Value`). `PostForm` is a part of `Main`. `Bar.NoScript` is the
  bar of an empty selection: a page with a script does not draw it (the bar
  comes with the first tick), a page without one always does.
  `TestSelectionForm`.
- **A thumbnail whose picture did not arrive keeps its icon** (Carrel's task
  148). `.tc-thumb__img` stands over the icon of the file's kind
  (`position:absolute`); `tc.js` takes away a picture whose address refused
  or whose file is no picture, and the tile shows the icon, as a file without
  a thumbnail does. With a script the icon is not drawn under a picture that
  arrived (a transparent one shows the thumbnail's colour, not the icon);
  without a script the icon stays under the picture. Nothing to do beyond
  passing `Icon` (above); `TestThumbPictureGone`.
- **A group of fields that is a form of its own** (Carrel's task 149).
  `FieldSet.Action` (with `Method`, `Fields`) prints the group as a form: a
  page of settings holds several — a name with its Save, an address with its
  Change — and each button sends its own group's fields. Such a group stands
  in a `Form` without an `Action`; inside a form that sends, the library
  prints it as a plain group — a form holds no form.
- **A pick that sends looks like every other pick** (Carrel's task 149).
  `Pick` with `Name` and `Options` prints the field `tc-in is-pick` with the
  browser's list inside it (`select.tc-in__list`): the field draws the frame
  and the arrow, the list only the words — before, the `select` drew an arrow
  of its own and stood 2 px lower than a pick that is a button. `Pick.Mark`
  is drawn in it too. `shape` drops the list and does not look into such a
  pick. A check that is a checkbox looks like a check when it stands in a
  field of a form. Nothing changes in the consumer's code.
- **An end follows its start** (Carrel's task 152, the owner's decision of
  11.10). `Input.Follows` names the field of the same form an end's date or
  time follows; when the start is changed, the library's script moves the
  end by as much, so the length stays: 10:00–11:00 with the start set to
  10:30 ends at 11:30. `Input.Day` names the date field that holds a time's
  day — an end carried over midnight takes its date along. An empty end
  stays empty, a field switched off (`Check.Disables`, «All day») is not
  touched, an end set by hand follows with its new length. Without a script
  nothing moves: what is on the screen is sent. The frames name the fields
  of an event's form: `start.date`, `start.time`, `end.date`, `end.time`
  (s-agenda-3, s-agenda-15, states-17) and `date`, `start`, `end` of the
  quick sheet (s-notes-9).
- **An upload without a script, wider** (Carrel's task 119). `FileField`
  has `Many` (several files at once), `Folder` (the system's dialog for a
  folder; every file goes named by its path inside it — the browser does
  that itself) and `Size`: with `Size: "size"` the library's script writes a
  hidden «size.<name>» with the length in bytes before each picked file, so
  that a server passing the body on as a stream knows it; without a script
  the fields are not sent. The words of a drop place that pick a file are a
  new part of `Empty`, `FilePick{Label, File}` — the label of its file
  input; its files wait for the button of the form they stand in. Frames:
  the Files and Folder of the Upload menu (s-files-8) are labels of file
  inputs with a no-script «Upload» under them, «choose on the computer» of
  the upload sheet (s-notes-11) is a `FilePick`; nothing looks different.
  Two file inputs of one name in one form: the one left empty goes as a part
  without a file name, as a browser sends it — the server skips it.
- **The picked files of an upload sheet are shown by name** (the owner's
  decisions of 09.10). `FilePick.Row` (a `*Line`) with `File.List` makes
  each picked file a row of the list — its name, its size, a Remove of its
  own (every file gets an input of its own, so a Remove takes one file
  back); a picked folder is one row, «сканы/ 14 files, 38 MB». The list is
  a `Line` inside the `Empty` (a new part of it), under the words of the
  place, and takes no room while it is empty. s-notes-11 draws three rows;
  the sheet's button says «Upload», without a number (s-notes-11,
  s-files-19). Without a script the browser's own field names the files.
- **A form with a file field of its own sends files.** `Form`, `Doc`,
  `Card` and `Sheet` with `Action` print `enctype="multipart/form-data"` when
  a `FileField` without `Form` stands in them: an edit's Save (the owner's
  decision of 02.10) and an upload sheet sent only the files' names before,
  script or none. Nothing to do on the consumer's side.
- **The table of published links breaks no word and no date** (the
  owner's decisions of 09.10). A `Table` with `Variant: TablePacked` keeps
  its mono cells on one line; a column that does not fit leaves, and one
  that has left does not come back while the rail's width is the same.
  Kind leaves first (`TableCol.Rank: Rank1`, at 880 px of the panel),
  Created after it (`Rank3`, at 752 px), Last opened as a secondary column
  does (700 px), Expires on the phone (s-publish-1, -17). When the rail
  folds the panel is wider, and the columns it has room for stand again.
- **Published links on an instance without mail** (the rest of Carrel's
  task 121): a new fragment s-publish-17 — the rows have Copy and Revoke and
  no Send, neither as a button nor in «⋯»; on a narrow width «⋯» holds
  Revoke alone. 426 frames.
- **Both reminders where the event is read** (the rest of Carrel's task
  106): the panel (s-agenda-1) and the card (s-agenda-2) say «10 minutes
  before · 1 day before», as the edit does (s-agenda-3, states-17); in the
  panel Reminders stand after Repeats, as on the card.
- **The numbers of the frames are the numbers of the rows drawn** (Carrel's
  task 143; the owner's decision of 08.10, Р-47). The count of a rubric is
  the number of rows under it; segments, subtitle, poll bar and rail are
  sums of what is drawn. One set for every frame, taken from the combined
  views: 5 contacts, 6 events of the week, 6 tasks (5 open and one done), 3 notes; the navigation
  says Contacts 5, Calendar 6, Tasks 2 (overdue), Notes 3. A frame of one
  collection draws its rows of the combined view (s-tasks-15, s-notes-31); a
  frame with a list under a sheet draws the whole screen: its subtitle, its
  rail and every row. Every kind counts in its subtitle, a folder of files
  too («5 items»), and a read-only collection counts before its reason.
- **Frames of one collection, the calendar subtitle, the poll bar**
  (Carrel's tasks 137–142; Р-48). A row in the list of one collection has no
  source and no column for it — `Single` of `TimeRow`, `FaceRow`,
  `ExcerptRow` and `DueRow` (`tc-row--single`) —, and a linked row names its sources in the hint of its badge; with
  nothing ticked the combined row of the rail is open and counts nothing
  (states-7, -18, -25); the subtitle of a calendar names the week, the
  zone and the count of events, as the other sections count (states-13,
  -15); a poll that has ended says how
  long ago at the right of its bar («2 min ago»); the head of one book has its
  mark (states-10, -11); Related to has the row a link is typed into.
- **The frames of search, duplicates and the conflict agree with the rest**
  (Carrel's task 136; the owner's decision of 08.10, Р-46). «Duplicates N»
  of the navigation, «N to look at» and the segments count the groups that
  wait for a decision — two, in every frame; what the lists draw as linked
  («Собрать чемодан», «Что почитать про CRDT») is linked on s-dup-1, with
  Unlink and the merge. The rail of the duplicates frames names every source
  of the four kinds, as the rail of search does, with the number of records
  compared, and its foot has the davloom mark; s-dup-6 (nothing loaded) is
  the page of duplicates, not a dialog. The second copy of Артём lies in
  «Архив», where a record can be deleted; the options of a value are named
  by the source of the record. s-find-1 has the navigation of every other
  frame. The tablet frame of the conflict (s-conflict-2) draws the same four
  properties and four buttons as s-conflict-1. s-dup-4, -5, -6 name their
  address (`data-screen`, `data-part`). A new frame, s-io-14: the report of
  an import that skipped records of another kind has a fourth card,
  «Skipped», of the look of «Failed» without the alarm. 425 frames.
- **The files frames agree with each other, and the missing ones are drawn**
  (Carrel's task 121; the owner's decision of 08.10, Р-44). The selection
  bar stands above the filter bar and replaces nothing; every folder frame
  has the same head (s-files-2, -4, -5, dark-3 follow s-files-1). s-files-1
  names the address of a folder, `/app/files/{account}/{col}/{path…}`, and
  sorts by name. New state frames: an empty folder, a storage whose server
  did not answer, a read-only storage, no storage ticked (states-32 … -35).
  New fragments: the sheets of Rename, Delete and New folder of a folder
  (s-files-21 … -23) and a refused row of the queue (s-files-24). 424
  frames.
- **The frames of tasks and notes agree with each other** (Carrel's tasks
  105, 127; the owner's decisions of 08.10, Р-42, Р-43). The card of a task
  that is not overdue says nothing of it, and its due date is in the full
  form («12 August 2026, 18:00»), in the panel too; the panel has the card's
  order of fields and sections and a Source block. A task just ticked stays
  under Open, struck through, until the list is read again. A due date
  without a time has All day ticked and both times disabled; an empty time
  is a field of time. The words of the task's edit are English. The source
  of the note is the text its reading draws; Reading width has «The same
  note in other notebooks»; the name of an unknown property is monospaced
  everywhere. The search by typing in Related to (s-notes-4) is the picture
  of the live layer, with the rows of the field's list (s-tasks-10).
- **The states of tasks and notes are frames** (Carrel's tasks 97, 124), the
  same set contacts have. Tasks: nothing ticked, the poll running, a source
  refused, an empty list, a list whose server did not answer, a read-only
  list, the loss report of an edit (states-18 … -24); a new task (s-tasks-14)
  and one list (s-tasks-15). Notes: the same seven (states-25 … -31), a new
  note (s-notes-30) and one notebook (s-notes-31). s-tasks-1 and s-notes-1
  now name the combined view's address (`/app/tasks`, `/app/notes`); the
  address of a collection belongs to the frames of one list and one
  notebook. One summary of a poll for every section: «N failed» with its
  menu (Р-31). In a read-only list the row's tick is a sign, not a button:
  `DueRow` prints `Tick.Sign` (state `sign`), and the generator reads it
  from a `span.tc-tick` in the row.
- **The calendar's state frames are drawn like its list** (Carrel's task 108
  and the rest of 106). states-3, -6, -13, -14 and -15 have the counts of the
  navigation and the buttons of the rail's foot; states-13 draws two days.
  One calendar has a frame of its own, s-agenda-17, and s-agenda-1 names
  `/app/calendar` only. One way for every section: a view with nothing
  ticked has no tags in the rail and a disabled Refresh («Nothing is ticked
  to ask.» — states-7, -14, -18, -25); while a poll runs Refresh is disabled
  («The poll is still running.») and the poll's bar says how long it has
  been running (states-3, -8, -13, -19, -26); a poll that ended with a
  refusal says under the list what is missing (s-agenda-1, as states-6 and
  -9). A source whose server did not answer has the sign in the rail, not a
  dash, in the frames of one collection too (states-11, -22, -29). The
  poll's meter is red only when the poll has ended with a refusal, and green
  while it runs (the owner's decision of 07.10, Р-40: s-agenda-1,
  s-contacts-7, dark-1, phone-1, mentions-2, s-mailbox-21). A new note that
  is still empty shows no draft (Р-41). 416 frames.
- **`Button.Copied`** (the owner's decision of 07.10, Р-36): what a Copy
  button says for a moment after it has copied (`data-tc-copied`). Its right
  edge stays where it is and the button grows or shrinks to the left, with
  its usual padding, so the buttons after it do not move (s-publish-1,
  s-backup-3).
- **A file's name and a tile are links, a tile is picked without a script**
  (Carrel's task 114). `TableFile.Href` makes the name of a table's row the
  link — a folder opens, a file is chosen; the row is not a link, its tick
  and buttons are controls of their own, so `TableRow` has no `Href`.
  `Tile.Href` does the same for a tile's name; `Tile.Send` makes the tile a
  checkbox of the form it stands in (`tc-tile__box`, over the whole tile and
  with no look of its own: the picked tile is the outlined one). Drawn in
  s-files-1, -5, -6 and s-publish-1; the shape check takes a screen with or
  without them. Under a finger the link of a table's name has its 44 px zone
  in its own padding, and a tile's long name wraps. `TestTilePick`.
- **A tree's third level and a node that sends its value** (Carrel's task
  116). `TreeLevel58` is three levels down. `TreeRow.Send` makes a node a
  radio button of the form the tree stands in (`tc-tree__box`): the chosen
  folder goes under the tree's name with the sheet's button, without a
  script; the tree's `On` is the checked one, an unavailable node is
  disabled. Drawn in the sheet of s-files-4 (`to`), which is a form there
  with «Move here» as its submit; the generator of screens reads a sheet's
  form and a submit button (`Sheet.Action`, `Button.Submit`). `TestTreeSend`.
- **Remove of a repeated field's row, in place** (Carrel's task 86):
  `Button.Removes` on the Remove of a phone, an address — with the library's
  script the row's fields are emptied and hidden where they stand, under
  their names, and the labels of the first row go to the row that is first
  now; without a script the button submits (`Submit`, `Name`, `Value`) as
  before. `TestRemoveRow`.
- **Copy** (Carrel's task 118): `Button.Copy` is the text the button puts on
  the clipboard, `Button.CopyFrom` the ID of the element whose text it puts
  there (an `Input`'s value, a `Diag`'s text) — with the library's script
  (`data-tc-copy`, `data-tc-copy-from`). Without a script a button is not
  shown and the text is selected by hand; a button with `Href` stays a link
  to the page that shows the text. Drawn on the published links (s-publish-1)
  and on the script of a backup job (s-backup-3). `TestCopy`.
- **An import's file field as drawn** (Carrel's task 133): Browse is the
  label of the file input (`Button.File`); `FileField.Shows` names the line
  that shows the picked file's name — `Input{Picked: true, ID}`, read-only,
  not shown without a script, where the file input shows the name itself;
  Preview file sends the form (s-io-3, -4).
- **A group of duplicates is a form** (Carrel's task 131): `Card.Action`,
  `Method`, `Fields`, `ID`; a property's values are radios
  (`choice.FN`, `choice.ORG`, `choice.PHOTO`), the card's buttons send `op`
  (link, not, merge, keep, unlink, offer). `Field.Action` makes a field a GET
  form of its own: the threshold in the head goes with Enter (s-dup-1, -3,
  -7).
- **A conflict is a form** (Carrel's task 132): `Diff.Action`, `Method`,
  `Fields`, `ID`; a property's versions are radios of its own name
  (`Check.Radio`, `choice.<n>`), «Write this combination» sends them, «Use
  only my edit» and «Use only the server version» send `all`, Cancel is a
  link back (s-conflict-1…5).
- **`SizeRow.Href`** (Carrel's task 129): a found file's name is the link.
- **A sheet is a form** (Carrel's task 134): `Sheet.Action`, `Method`,
  `Fields`, `ID`, as `Form` and `Doc` have them.
- **`Swatches.Name`** (Carrel's task 135): the colours are radio buttons of
  the form, the chosen one sends its colour without a script; the checked
  one is outlined. The new collection's sheets draw it so (s-profile-5, -7,
  s-mailbox-22, -27).
- **A note's text in its panel** (Carrel's task 125): `Sec` takes the parts
  of a document's text — `DocHeading`, `DocBullets`, `DocQuote`, `DocCode`,
  `DocTable`. The new frame «Заметка в панели · части текста» (s-notes-28).
- **A drop in an edit's bar** (Carrel's task 126): `Docbar` takes a `Drop` —
  Attach of a note's edit is the menu it is in the column (s-notes-4, -20).
- **A note without a script** (Carrel's task 104, the owner's decisions of
  02.10). `Docbar.Edit`: an edit's bar keeps only Save and Cancel without a
  script — an edit is not reloaded midway; the marks, Auto-save, the draft,
  Ctrl+S, the width, Focus and «⋯» need the page's script. On the reading
  page Source, the width of the text, the side column and «⋯» are links
  (`?view=source`, `?width=…`, `?side=0`; «⋯» opens its menu of links):
  `Button.Swap` and `Seg.Swap` name what the live layer takes from the page
  at the address and puts in place — no reload, the address in the history
  (`Main.ID` names the column); without the layer the link goes there. The
  source of a note is its own page — the new frame «Заметка · исходник»
  (s-notes-27), Rendered goes back. `TestSwap`.
- **A file of an edit goes with its Save** (the owner's decisions of 02.10:
  an edit is not reloaded midway, the same in every edit). `FileField.List`
  makes the field one of its form's: the library's script shows a picked file
  as the menu item's `Row` (`PickedRow`: a `DocAttach` or a `Line`,
  «{name}» and «{size}» in its words) at the end of the list named, the input
  inside it; `Button.Removes` takes the row back before Save; `Dim.Empty`
  («Nothing attached») hides while a row stands. `FileField.Preview` shows a
  picked photo at once in its `Photo.ID`. `Runbar.Progress` is the Save line
  of a form that sends files: hidden until then, the live layer moves its
  meter and writes «1.5 of 3.0 MB» (`Saving` draws it so). `Line.ID`,
  `DocmetaSec.ID` name the lists, `DocAttach.Remove` is Remove on an edit. A
  picked file sends its form at once only where the field names a `Form` or a
  `To`. The note's, the event's, the new event's, the task's and the
  contact's edits draw it so (s-notes-20, s-agenda-3, -15, s-tasks-3,
  s-contacts-3, -4, -15, -27, states-12); the new frame «Правка заметки ·
  Save с файлом» (s-notes-26) draws a picked file and the Save line.
  `TestEditFile`.
- **A menu item and a file that send the form they stand in** (Carrel's task
  122): `MenuItem.To` is a submit button of the surrounding form with its own
  address (`formaction`), as `Button.To`; `FileField.To`, without `Form`,
  sends the surrounding form there when a file is picked. Attach on a new
  event and in a task's edit keeps what is typed (s-agenda-15, s-tasks-3).
  `TestFileInForm`.
- **`ColumnRow.Href` with a tick** (Carrel's task 123): the row is not a
  link, its name is, as `DueRow`'s title.
- **A row's «⋯» with its menu** (Carrel's task 117): `TableCell.Drop` — the
  row's «⋯» and the twins of its is-2nd buttons, a details that opens without
  a script; `Drop.Row` draws its trigger as the row's «⋯». The rows of shared
  links (s-publish-1, s-mail-4) carry Send and Revoke under it.
- **A link option of a segment keeps its icon** (Carrel's task 115): an
  option with `Href` prints `Icon`, and `Title` names an option without
  words — the folder's List / Tiles (s-files-1, -5, dark-3).
- **A document is a form in its edit** (Carrel's task 98): `Doc.Action`,
  `Method`, `Fields` and `ID`, as `Form` has them. A note's edit sends its
  fields without a script: the title (`summary`), the notebook, the date and
  time, the tags, the text (`description`) — names the frames give and the
  server may change; Save in the edit bar is the form's submit (`Form`),
  Cancel a link back to the note (s-notes-4, -20).
- **A pick of collections with their marks** (Carrel's task 103): `Pick.Mark`
  shows the chosen collection's colour; a radio `MenuItem` prints its `Mark`,
  and with a `Reason` it is shown and switched off, the reason in its hint —
  a read-only notebook with its «ro» badge (`Trail`). Choosing an item, the
  core script carries its mark into the pick. The Notebook field of a note's
  edit is drawn so (s-notes-4, -20; its list s-notes-15, desktop-17).
- **Parts of a document's text** (Carrel's task 102): `Em` (italic words),
  `DocQuote` (a quotation), `DocCode` (a block of code as written; a long
  line scrolls inside it), and a nested list — a `DocBullets` item's parts
  may end with a `DocBullets`. A link and inline code (`Mono` chip) stand in
  a paragraph and in an item as before. The new frame «Заметка · части
  текста» (s-notes-25) draws them all.
- **`SegFold.Phone`**: a short switch folds on a phone's bar only (640 px
  and less), where the timeline's folds with the bar's second rank (960).
- **`FilterH26W200`** (Carrel's task 100): a bar's filter of 200 px.
- **`Tick.Sign`** (Carrel's task 99): a task's tick in a text — a note's
  tasks, its side panel, a published page — is a sign that says «Done» or
  «Not done» and is not pressed; the note is changed in its edit. The frames
  of notes (s-notes-1, -2, -3, phone-7) and of a published page (s-publish-8)
  draw it so.
- **`TagAdd.Options`** (Carrel's task 96): the choices of a tag field's line
  on a page without a script — a `datalist`, the browser's own list as one
  types; a `ListOption`'s Value is sent (a record's UID, an address), its
  Label shown beside it. Where the field has `Suggest`, the live layer takes
  the list away so that it does not cover the menu.
- **`Check.Disables` takes several names** (Carrel's task 111), separated by
  spaces as classes are: All day switches off «start.time due.time». The
  core script switches off each.
- **A file without a script** (Carrel's task 109): `Button.File` and
  `MenuItem.File` (a `FileField`: the `PostForm`'s ID, the field's name,
  `Accept`, the input's accessible name) make the control the label of a file input, which
  the browser opens without a script. With the library's script the picked
  file sends its form at once; without one the input shows what was picked
  and an «Upload» beside it — `Button.NoScript` / `MenuItem.NoScript`, shown
  only on a page without a script — sends it. «From this computer» in every
  Attach menu and the photo's Replace and Add are drawn so. `TestFilePick`.
- **`Range.Hidden`** (Carrel's task 107): the range's form sends the list's
  other choices with it — `view=range`, without which a server shows the
  week from From. The agenda's ranges carry it.
- **`DueRow.Href`** (Carrel's task 95): a task's row opens its record without
  a script — its title is the link, the row cannot be one with its own tick.
  The list of tasks draws it so (s-tasks-1).
- **Save as** (Carrel's task 94): `Button.SaveAs` and `SaveFrom` save the
  text of the element named (a `Diag` with its new `ID`) as a file, made in
  the page by the library's script — a secret shown once has no address to
  download from. Without a script the button is not shown; Copy and
  selecting the text remain. `TestSaveAs`.
- **Refused sources** (the owner's decision of 02.10: the fight with the
  length of a line is lost): a `TickRow` with `Error` and no `Meta` shows an
  error mark, its reason (`MetaTitle`) in the mark's hint (the icon `alert`);
  the runbar says «N failed» in red — a closed `Drop` (now a `RunbarPart`)
  whose menu has a `MenuRefusal` per source: its mark, what was asked and
  what the server answered, in full, and its Retry, with Retry all below
  (`states-16`). Every frame with a refusal is drawn so; the agenda's «⋯»
  is closed on the screen and its menu has a frame of its own (s-agenda-16).
  The ladder does not measure what a closed details hides.
- **A bar's checks under «⋯»** (Carrel's task 82): `CheckForm` is a
  `MenuPart` — in a menu its checks are menu items with tick boxes, in the same
  GET form, so the twin of Compact rows works without a script too (Apply
  shows only then). «⋯» with its menu is `Drop{More: true}` among a bar's
  `Items`; `Bar.More` stays a bare button. `TestCheckFormInMenu`.
- **A tag field that is a part of its form, with suggestions** (Carrel's task
  77, the attendees of an event): `TagInput.Columns` (tags in columns across
  the width), `TagInput.Name` (each tag with a `Value` sends it; a tag with a
  cross is a checkbox the cross unticks, so without a script too), `TagInput.Add`
  (`TagAdd`: the typing line, sent as `Name` without a script; `Suggest` — the
  library's script asks `Suggest?q=…` and shows the menu items the server
  answers). `MenuItem.Adds` is the tag an item adds (a template), and
  `MenuItems` prints the items of that answer. `Tag` gains `Address` (an
  unknown person, monospaced), `Aside` (their reply), `Title`, `Name`, `Value`
  and `Fixed` (a pill without a cross: the organiser). `SecList.Columns` sets a
  section's list in columns. Enter in the typing line chooses and never sends
  the edit; a cross and Backspace take a tag away. `TestAttendeesSuggest`.
- **The live layer, `tc-live.js`** (the owner's decision of 02.10, L-1): the
  library's second script, loaded after `tc.js` when a page asks for it
  (`page.Document.Live`; `page.At` names `js/tc-live.js`, `page.Static` and
  `typecase.LiveJS` serve it). It asks only addresses the markup names, of
  the page's own site, and puts the HTML of the library's widgets where the
  markup says; the core, `tc.js`, has no network and no storage, and opens
  `window.__tc.menuOf`, `showMenu` and `closeMenu` to the layer. A tag field's
  suggestions live there; without it the typed words go with the form.

- **`Query.Placeholder`**: what an empty search line says — «Search
  everything…», «Filter the loaded cards» — is a placeholder, not a value; the
  frames draw it so, in the colour of the text (Carrel's task 69).
- **`Tag.Href`**: a tag goes somewhere — the list filtered by it (task 63); in
  a row that is a link itself a tag stays words, not a link inside a link.
- **`HeadPhoto.Src`**: the photo in a card's head and a timeline's panel (task
  64); for `shape` the picture and the placeholder's `is-ph` are data.
- **`Filter.Hidden`**: what a filter's form sends with the query — the list's
  sort, tag, kind (task 65).
- **`MenuItem.Form`**: an item that submits a `PostForm` by its name — Sign out
  (task 68; drawn so in account-5).
- **A closed menu opens without a script** (task 60, U-8): a closed `Drop` or
  `TopDrop` is a `details` whose `summary` is its button, in the button's
  look; the library's script opens and closes the details and closes it on Esc
  and on a click elsewhere. The card's «⋯» is drawn so (s-contacts-6, -7).
  `TestClosedMenuWithoutScripts`. The ladders see what a closed details hides
  (`checkVisibility`).
- **`CheckForm`** (task 67): a bar's check that is a GET form — «Compact rows»
  goes with the list's other choices; with the library's script a change
  sends it, on a page without a script its Apply does. A page says it has no
  script until `tc.js` runs (`data-tc-noscript` on `html`, printed by `page`).
  Drawn so in the lists' bars (s-contacts-1, -2, s-agenda-1);
  `TestCheckFormSubmits`. `CheckForm.Checks` holds several checks of one form:
  the agenda's «Events · Tasks due · Notes» (s-agenda-1, task 74); with
  `Tertiary` they go under «⋯» on a phone. A bar takes checks only so.
- **`Seg.Multiple`**: a switch with a Name may be a row of checkboxes —
  several options at once, each ticked one sent: the days of a repeat
  (s-agenda-3, task 72).
- **A pick with a menu to search and choose from** (task 73): a time zone
  is a closed `Drop` with a `Pick` and a menu of radio items (`MenuItem` with
  a `Name` and no `Tick`) — a search line, «As sent», rules. Without a script
  the chosen zone goes with the edit form and the search line is not shown;
  with the library's script the search keeps what matches and the choice
  becomes the pick's value (s-agenda-3, -8). `TestZonePick`. A record may
  stand on a closed menu's `summary`: it is its trigger button.
- **`shape` is strict about what is missing** (task 76): a run of the screen
  is no shorter than the frame's, and each element the frame draws in it is
  on the screen but for its states — a head's Print, a card's Related, a
  field's time and badge missing is red. A repeat of rows keeps any length;
  repeated as often as the frame draws it or more, it has each part the frame
  draws. **A client whose list shows fewer rows than its frame now fails:
  give the check the frame's records.**
- **`shape` skips templates** (task 75): a `<template>`'s contents are not
  compared. (A repeat's blocks of another kind — the empty days of a week —
  are not needed: by the owner's decision of 02.10 a day without events is
  not printed.)
- **`Drop` in a rubric and in a line** (task 71): Attach with its arrow —
  «From this computer», «From Files» — stands in the Attachments rubric of a
  card and in the attachments line of an edit, closed, opening without a
  script (s-agenda-1, -2, -3, s-tasks-1..3, s-notes-1..3, -20, s-io-2).
- **`shape` compares a repeating row of rows at any length** (task 61): a
  month and its records, a term and its value — the screen may repeat the
  block any number of times, each like one of the frame's; terms and values
  repeat by tag, so `dd` and `dd.tc-mono` are one value. `TestRepeatsAnyLength`.
- Frames drawn whole (task 62, the owner's decision of 01.10): the read-only
  book with its bar and groups (states-4), the edit without a photo and the new
  contact with their whole forms (s-contacts-4, -27); the second phone's label
  is a pick, as the form sends (s-contacts-3). 373 frames.
- The states of contacts have their frames (task 66), in the chapter of
  states: no book ticked (states-7), the poll still running (states-8), a source
  refused (states-9), an empty book (states-10), a book whose server did not
  answer (states-11), an edit whose save lost properties (states-12). 379
  frames.

### Mockup

- **A note keeps what the server sent** (Carrel's task 253, accepted by the
  owner 03.11): the note's side column, above Source, has «Kept as the
  server sent them» folded with «Show 4» and the contact's hint — a second
  `DESCRIPTION` and three properties Carrel does not know (`s-notes-2`,
  `-3`, `-27`; not the guest's page `s-publish-4`).
- **A queue row refused for want of the size** says «server requires the
  file size» (Carrel's task 252, the owner's words; `s-files-2`, `-24`).
- **Tables as cards, the fourth group** (Р-66, accepted by the owner 03.11):
  the setup check and the DAV check with What happened (`s-admin-4`, `-29`;
  the card's name is the check or the operation, not its status mark), the
  audit log with Action (`s-admin-7`; the name is the time), the password
  history with What (`s-auth-15`; the name is the time). A card whose row
  begins with a tick or a status mark takes the next cell as its name.
- **Tables as cards, the third group** (Р-66, accepted by the owner 02.11):
  device setup with Type (`s-plan-4`, `-5`: where to paste the URL; Copy
  copies it), the backup history with Verified (`s-backup-1`, `-12`…`-14`;
  the card's name is the date), the restore match with Restore into
  (`s-backup-8`, `-20`: the field to choose), the restore report with What
  happened (`s-backup-18`).
- **One selection bar** (accepted by the owner 02.11): every selection bar
  ends with Select all and Clear selection, icon and label (a new icon
  `i-selectall`; `s-files-1`, `-2`, `-4`, `-5`, `s-backup-8`, `-20`,
  `s-mailbox-3`, `components-16`, `-17`, `dark-3`, `phone-6`). Where the
  labels go (narrow), the two icons stand on the count line, at its end;
  the actions are the second line. Up to a 1280 px screen the two are icons
  (with their hint): a mail bar with every action had no room for both words.
- **The notes list on paper** (Carrel's task 250, accepted by the owner
  02.11): a new frame `print-8`, like the tasks on paper (`print-4`): the
  title with the notebook, the notebooks in the snapshot, a «Date | Note»
  head, the month rubrics; a row is the date, then the title and under it,
  dim, the excerpt and the notebook («in 2 notebooks» for a linked note).
  The current view is printed: filter, order with its rubrics, columns —
  Tags, when on, join the line under the title. The foot counts the notes
  and the notebooks that answered. 570 frames.
- **The owner's card in the combined contact lists** (Carrel's task 251,
  accepted by the owner 02.11): «Михаил Енученко» (`mix@example.org`,
  nextcloud · Контакты, no phone, no organisation) stands as group «М»
  after «Л» in every combined list — `s-contacts-1`, `-2`, `-35`,
  `s-notes-9`, `s-find-2`, `states-8`, `-9`, `phone-4` and on paper
  `print-2`. The counts are one card higher: Contacts in the rail 8, the
  combined address book 10, «Контакты» 6, «7 cards» in heads, poll bars
  and the paper foot, the loaded counter by Compact rows 7.
- **Import previews as cards** (Р-66, accepted by the owner 02.11): Result is
  the main column of every import table (`s-io-5`…`-8`, `-19`). An import of
  one file — contacts, events, tasks — has no Source column: every row said
  the same file the head already names; the notes import, the one that takes
  many files, keeps File first. The contacts preview shows Email, not UID: a
  person tells two «Анна Ковалёва» apart by the address (Name | Email |
  Result). On a phone a contact, an event and a task open with their name,
  a note with its file.
- **Tables as cards, the second group** (Р-66, accepted by the owner 02.11):
  sessions with Last seen (`account-1`, `-15`, `s-mail-9`), published
  links with Expires (`s-publish-1`, `-17`, `-20`, `-21`, `s-mail-4`),
  invitations with Expires (`s-admin-1`; a link invitation's address reads
  «by link», not «—»), escrow with Deposited (`s-admin-6`, `-30`; «no», not
  «not deposited»). A card keeps the button before «⋯» (Copy link, a file's
  Download) and «⋯»; where «⋯» is a fold (`tc-acts__more`), only what folds
  into it (`is-2nd`) goes — a link's Copy stays.
  On a phone a field with `flex:1` in a line that turned into a column is
  38 px, not squeezed to its text (`tc-in--fill`: `account-1`, `-8`…`-10`,
  `attachments-1`, `-5`, `-6`, `s-admin-3`, `-32`, `-33`, `s-backup-2`,
  `-19`, `s-davloom-1`, `-6`…`-9`, `s-plan-2`, `desktop-13`). The selection
  bar puts its count on the first line and every action on the second — the
  separator turns into a line break — instead of one button falling down
  alone (the owner's decision 02.11; `s-files-1`, `-2`, `-4`, `-5`,
  `s-backup-8`, `-20`, `s-mailbox-3`, `components-16`, `-17`, `dark-3`).
  At 1351 and 820 only the two texts differ.
- **On a phone a table is cards with its main column** (Р-66), the first
  group, accepted by the owner: users with Role (`s-admin-2`, `-22`, `-24`,
  `-34`), files with Changed (`s-files-1`, `-2`, `-4`, `states-34`,
  `s-publish-5`, `-12`, `-30`…`-32`, `s-mail-6`, `-13`), connections with
  Changed on the server (`s-profile-1` and its family; What is in it in
  `s-profile-7`…`-9`), devices with What it can see and do (`s-davloom-1`,
  `-7`…`-9`). A card's first line is the name — the cell with a file, or
  the first one not empty and not a tick — and the row's last button,
  «⋯»; a row with one button and no «⋯» (Revoke) keeps it. Declare the
  column with `TableCol.Main`.
  On a phone the floating upload queue stands under the list, not over its
  last row (`s-files-2`, `-3`, `-24`; folded, it is a bar
  as wide); in a form that turned into a column a
  field's basis no longer becomes its height (`s-davloom-1`, `-6`…`-9` had
  ~460 px empty under the address); a switch in a table row is 38 px, as
  the row's buttons (Р-65). At 1351 and 820 no frame differs.
- **A person's timeline filters by word** (Р-63: a filter by name does not
  care which record): «Filter by word» first in its bar, as in the agenda
  (`s-contacts-7`). A bar with a filter and a date range wraps where both
  do not fit — neither goes under «⋯» (`.tc-bar:has(> .tc-in):has(>
  .tc-range)`, Р-65): at 1191, beside the card, the range goes to a
  second line.
  Its poll bar says what was polled, what failed and when — no «matched by»
  with the card's addresses (Р-77: the column «why» names the field of each
  row); `s-mailbox-21` too.
- **Print in every list's head** (Р-63: printing does not care which
  record): the notes list had New note, Import and Export only —
  `s-notes-1`, `-31`, `-33`, `states-25`…`-29`, `phone-8`, `-9`; the lists
  under an import drop and the dark list too (`s-io-1`, `-2`, `dark-2`).
- **The 580 column holds in every frame of a screen** (Carrel's task 246,
  Р-82): a field's hint and a message stand in it on the page, in a dialog
  and in a side panel, never in a card or a sheet — 25 frames brought to
  the rule (`account-1`, `-2`, `appearance-1`, `s-admin-1`, `-3`, `-30`,
  `s-backup-2`, `-7`, `-8`, `-15`, `-16`, `s-davloom-2`, `-3`, `-6`,
  `s-profile-2`, `-10`, `-11`, `-14`, `-16`, `-22`…`-24`, `-31`, `-32`,
  `s-mail-4`); 19 of them look the same to the pixel. The link letter's
  Send says «Sending…» too (`s-mail-4`).
- **Frames caught up with their neighbours** (Carrel's task 244):
  `states-24` is the task's edit `s-tasks-3` with its loss line — the
  repeat's hint under Advanced RRULE, «Kept as the server sent them» with
  SEQUENCE only, Read «just now · changed on server just now»;
  `s-contacts-16` is Anna's card `s-contacts-6` with the rewrite's message,
  the work address the new one; Anna's edits (`s-contacts-3`, `-4`,
  `states-12`) have her second address and her work address, as the card.

- The archive's line of a deleted collection is short in all its cases
  (Carrel's task 240): «No backup job copies it: there is no archive to
  restore it from.», «A backup job copies it, but none has finished whole
  yet: …»; a calendar's restore makes «a new calendar» (`s-profile-9`).

- **The page under the letter is the combined address book** (Carrel's
  task 241): the rail, head and rows' sources of `s-contacts-1` in
  `s-mail-1`, `-5`, `-6`, `-10`, `-12`, `-13`, `-16`…`-18` — no calendars,
  no baikal; the head «Combined address book · 3 of 4 books · 2 accounts ·
  2 selected».

- The letter's Send says «Sending…» while the letter goes (`s-mail-1`,
  `-14`, `-17`; the drawn state `s-mail-19`, Carrel's task 242).

- **Dates in words, the weekday apart** (Carrel's task 243; the owner,
  28.10): the import preview writes dates as the card does — a series'
  first start with its rule («Monday, 10 August 2026, 09:00 · weekly until
  28 December»), a span of whole days by its first and last day («13 – 26
  July 2026 · all day»), a foreign zone by the person's time, the event's
  time and zone in the mark (`s-io-6`, `-7`, `-8`), in the plain face. A
  weekday stands before the date with a comma («Monday, 10 August»), and a
  series ends in words on the card («until 28 December»).

- **Settings and administration read in one column** (the owner, 27.10):
  on 74 frames of the settings and administration pages fields' hints and
  messages stand in a column of 580 px; Display name, Address, Folder in
  it, Written as, Public address are as wide as the column; «Separate from
  the sign-in password.» is the password's hint (`s-backup-2`, `-19`).
  «Signed in as» (`desktop-13`) and «Address for the client» (davloom)
  name their fields for a screen reader.

- **What follows Send** (the owner, 27.10; Carrel's task 194 p. 7): all
  accepted — the sheet closes, the screen under it says «The letter went to
  2 recipients.» (`s-mail-16`); all refused — the letter as it was, the
  relay's words first under To (`s-mail-17`); some refused — a sheet of who
  got it and who was refused, «Send again to the refused» (`s-mail-18`). The
  line under To reads «Up to 10 recipients. Each sees the others'
  addresses.».

- **The letter's Cancel asks when a text is typed** (the owner, 26.10;
  Carrel's task 194 p. 7): Cancel sends the form (`op=close`); without a
  text the sheet closes and the draft is forgotten, with one — new
  `s-mail-15` «Close the letter?»: «Close and lose the text» (`op=discard`)
  or «Keep writing» with Esc. The letter's drafts wait for the mail section.

- **Every sheet names its keys** (Р-81, the owner 26.10): the main button
  «Enter», or «Ctrl+Enter» in a sheet with a text (a letter, a reply, a
  quick note); Cancel «Esc» at the end of the foot — 31 sheets mended. A
  button with its key does not shrink in the foot (`.tc-sheet__foot >
  .tc-line, > .tc-btn, > .tc-kbd{flex:none}`). The keys sheet (`account-6`)
  says Ctrl+Enter saves or sends a sheet with text.

- **The letter's Preview shows it as the recipients get it** (the owner,
  26.10; Carrel's task 194 p. 7): new `s-mail-14` — From, Reply-To, To,
  Subject, the text with its signature, files and the block of links; Send,
  Edit, Cancel. Reply-To carries the name as From does; the signature
  stands after «-- » (RFC 3676) in `s-mail-7`, `-8`, `-14`. A tag with an
  icon keeps it on the name's line (`.tc-tag:has(> svg)`).

- **A file chosen from the computer can be dropped too** (Р-70; accepted by
  the owner 26.10): new «Бросок там, где файл выбирают» — a table of which
  page carries which drop sheet (`Scrim{Drop, Closed}`) and the sheets
  themselves: a record on its own page «Drop to attach» (`s-io-20`), an edit
  «Attached when you save» (`s-io-21`), a contact's photo (`s-io-22`), a
  restore from this computer (`s-io-23`); the import wizard's first step
  carries the list's «Drop to import».

- **Explanations on pass-through screens say only what is important and
  not obvious** (Р-77; accepted by the owner 26.10): about 50 sheets and
  dialogs lose the mechanics and the captain's lines and keep no counts
  the server would have to compute (the owner, 26.10). Deletions say
  «Carrel keeps no copy: this cannot be undone.»; a folder's adds
  «Everything in it goes too.» (new `s-files-27`, «Файлы · удаление
  папки»); the import wizard's steps lose their reassurances.
- **Restore asks what import asks** (Р-63; the owner, 26.10): records of the
  archive already here and different — «Skip them» by default, «Replace
  them», «Import as copies», above the table (`s-backup-8`, `-20`); the
  report counts what was skipped (`s-backup-18`). «Restoring is an import:
  records made since the archive stay, nothing is deleted.»

- «Keyboard shortcuts» in the account's menu (`account-5`, `desktop-2`) is
  a link to `#keys` pressed by «?» and shown only with a script: without
  one the keys do not work either. The keys sheet (`account-6`) is
  `#keys` (Carrel's task 239).

- The guest's address book (`s-publish-7`) draws no row in bold: a guest's
  rows are links, none is open (Carrel's task 238).

- **Paper and the offline page** (Carrel's task 228; accepted by the owner
  25.10–26.10): new «Остальное на бумаге» — the contacts list, a contact's
  card with the properties kept as the server sent them, tasks of the
  current filter, a note in reading width, a folder, one calendar's agenda
  with its own title (`print-2`…`-7`); tables name their columns
  (`desktop-16` too); names in regular weight, drawn ticks, one line at
  the bottom. The contacts list prints without photos (the owner, 26.10).
  `offline-1` — the installed app without a connection: the mark, the
  message and Try again. A struck-through done item's line runs through the
  middle of the small letters on screen too.

- **A person's timeline has no «attached to» row** (Carrel's task 237; the
  owner's decision of 25.10): a file inside a record the timeline already
  shows is not repeated; files come only by the person's name in the file's
  name. The counts follow (Timeline 10; August 2026 · 2; All 23, Files 3).
  «Согласовать обложку» lies in «Работа», done; «Ревью вёрстки сборника» in
  «Проекты» on 29.07 — the task rails, the list's counts and the duplicates'
  sources count them.

- **The tasks of «Дела» as the set holds them** (Carrel's task 236;
  accepted by the owner 25.10): the МФЦ task's panel and guest page carry
  its priority, reminder, tag and description, no place and no
  attachment, «overdue by 3 days» under the panel's title as on the card;
  «Забрать документы» is related on 11 Aug.

- **One name for a filter's field, `filter`, on every screen** (Carrel's
  task 231; accepted by the owner 25.10): the owner's calendar and the
  guest pages send `filter`, as the owner's lists do — `q` stays the
  search's. The filtered calendar's series row says «4 attendees» and
  «repeats weekly», as in the week; a filter with a word has its cross
  (`s-agenda-31`, `s-publish-39`).

- **The quick sheet without its collection; «Custom…» in the quick sheets;
  the crop with its picture** (Carrel's task 233; accepted by the owner
  25.10). When no collection is remembered, the quick sheet shows its
  fields at once, «into» without a value («Write into»), and «Save note»
  waits for the choice (`ButtonWaits`, required radios) — `s-notes-34`.
  The quick event's and task's reminder list is the full form's, its
  «Custom…» count and unit hidden in the same row. `s-contacts-5` draws the
  server's picture. Rails count Anna's copy in Архив everywhere.

- **The hidden zones stand in the shell's scheme only** (Carrel's task
  235): the run's line and the zone of messages, waiting hidden, moved from
  `s-agenda-17` to `shell-1`, beside the sheet's place — a screen's frame
  draws none, the consumer's document puts them. `states-11` draws its
  message and technical block as one zone of messages. The look of every
  frame is the same.

- **Guest pages against the owner's frames** (Carrel's task 234; accepted
  by the owner 25.10). «Личный» holds what the owner's calendar does — the
  weekly sync-up on Mondays (4 attendees), «Забрать документы» on Tuesday
  11.08 at 18:00; «Команда» holds Anna and Artём; «Записи» two notes; the
  note ends with its last paragraph; the МФЦ task is not begun. A guest sees
  an attendee by the name in CN, by the address only when there is none.

- **The linked record and duplicates against their neighbours** (Carrel's
  task 232; accepted by the owner 25.10). Copies stand in the rail's
  order — by connection, then the collections' order in it — in a group's
  rows, its questions and «Where … is kept»; the contacts' linked panel has
  «Kept as the server sent them»; the card after a merge has the contact
  card's head, its timeline rows link to their records («related to this
  person»), as every contact card's do; the linked task card is drawn as
  the task card, the linked event repeats «until 2026-12-28», the linked
  note's column is Tags, Related, Attachments; «Заметка со встречи» is
  linked in its notebook too. New: the panel of a task and of a note that
  lie in one collection (`s-tasks-20`, `s-notes-33`). Anna's other books
  are Архив and Команда in three more frames.

- **Duplicates after note 117: the frames against the fixture set and
  each other** (Carrel's task 230; accepted by the owner 24.10). Anna's
  copies share one UID — «220 points · same UID, …»; every copy has its
  organisation; a variant says its value and collection everywhere; a
  field some copies lack is asked with «—» (Title, Address, Birthday,
  Note); Vera's three addresses; «Decided earlier» has Artём; the empty
  screen's refresh is off with its reason. **A linked record shows what
  any copy holds** (Р-78): the contact's panel and card share their
  fields, the task's and the note's panels are linked records as the
  book's (Р-63), the linked event, task, note and contact cards have their
  description, attendees, attachments, related, timeline — with the mark
  of the copy they come from; Source and Delete stay with each copy. The
  card after a merge is the whole contact card. Rails count what lies in a
  collection (Contacts 9, Архив 2, Команда 2; Дневник 3); the timeline's
  other books are Команда and Архив; `states-8`, `-9` draw the linked rows;
  the copy sheet says «Due, priority differ · due 12 Aug»;
  `release-checklist.md` is 3.1 KB everywhere.

- **Public pages: what the guest sees, the views without frames** (Carrel's
  task 196; accepted by the owner 24.10). A row's caption is what the link gives — «repeats weekly» and
  the status always, «3 attendees» with Attendees ticked, the owner's marks
  never; a row of a published list opens its record's page in the same
  link (`?uid=`), a folder's name its listing (`?in=`) or the file's page
  (`?file=`). The foot is the link's day only, on every page; the line
  under a title says the days or the count. Downloads are links; the filter is a GET form
  (`?q=`), the width switch is addresses
  (`?width=reading`), «How to mount» opens without a script. A folder
  counts «items»; over the ceiling only the page's listing is cut — the
  archive and the drive give the whole folder. New: the whole calendar by weeks (`s-publish-24`), an
  empty week, book and folder (`-25`, `-35`, `-36`), an event, a contact,
  a task (`-16`, `-26`, `-27`), a task list and a notebook (`-28`, `-29`),
  a folder in a folder and over the ceiling (`-30`, `-31`), «How to mount»
  open (`-32`), a PDF and a text over the ceiling (`-33`, `-34`), «The
  source did not answer» and «No longer there» (`-37`, `-38`).
- **A calendar filters by word, as the other kinds do** (Carrel's task
  196; the owner, 24.10): «Filter by word» first in the bar of the
  owner's calendar (`s-agenda-1`, `-17`, `states-3`, `-6`, `-13`, `-15`)
  and in the head of a public calendar (`s-publish-6`, `-24`). With a word
  the list is everything found in the calendar by days — in a range link,
  in the range — and ← Today →, the period and the range go
  (`s-agenda-31`, `s-publish-39`); a repeating event stands once, at its
  nearest time. A public page filters with the owner's words: «Filter by
  word», «Filter this folder». A note's Markdown heading is drawn as a
  heading of the text everywhere (`s-notes-2`, `-3`, `s-publish-4`, `-8`,
  `-23`, the phone's note).

- **The rewrite sheets in one view; the download of a link off by default**
  (Carrel's tasks 223, 229; accepted by the owner 23.10). `s-contacts-19`,
  `-20` are drawn as `s-contacts-15`: the replacement row with its choice,
  the groups by ticks, what stays as it was in one line over the button;
  with nothing to rewrite, «Rewrite 0 records» is off with its reason. The
  captions say the order of a group's records and the delete sheet's width.
  «Allow … download» is unticked in `s-publish-2`, `-4`, `-19`.

- **Sheets: the move's lines, «where» of the quick sheets, export, the
  password, the crop** (Carrel's task 197 pp. 3, 5; accepted by the owner
  23.10, Р-77). A sheet's note says only what matters and is not obvious —
  no note when there is nothing such: the move and copy sheets of records,
  collections and files are shortened or lose their note (`s-contacts-2`,
  `-17`, `s-files-4`, `-11`, `-12`, `-13`, `s-profile-7`, `-8`, `-18`,
  `-34`, `-35`). New frames: `s-agenda-29` (an event moved to another
  account: its attendees may get a cancellation), `s-contacts-37` (a copy
  under a new UID), `s-agenda-30` (the export sheet of a combined view),
  `s-profile-36` (a wrong password). «Where» of a quick sheet is a drop of
  radios `into`, the last used one first, a read-only one off (`s-notes-5`,
  `-8`, `-9`, `-10`, `-16`, `-17`, `-19`); a pushed drop in a sheet's head
  keeps its field at the right. The crop page is a form (`s-contacts-5`);
  the new card's form has no «Show 5 empty fields» (`s-contacts-27`).

- **Search after note 114** (Carrel's task 227; accepted by the owner
  23.10): the stopped search has its segments (`s-find-10`); one bar in
  every state — the words in two lengths, the menu of refusals after them,
  the time pushed right (`s-find-1`, `-4`, `-5`, `-9`, `-10`); Пётр Ковалёв
  counts for Команда (`s-find-1`, `-4`); the captions say when the line about
  files stands and that the suggestions come from the server.

- **Duplicates and the linked record** (Carrel's task 190; accepted by the
  owner 23.10). `s-dup-1`: every linked group is on the screen, one data
  set with the lists; single values are asked by a switch, many by «what to
  keep» ticks; an event's time is one question «When», a task's «Due»; the
  points add the weights (a name with its words swapped is the same name);
  the forms post to `/app/duplicates/decide`, the threshold too; a task's
  tick is a sign. `s-dup-3`, `-7`: the weights are a form. `s-dup-6`:
  «Decided earlier» without loaded lists. New `s-dup-8`: a merge that went
  through, with «Linked records». The linked record's card has Edit
  (→ «Which copy to edit?») and «⋯» — Print, Open in Duplicates, Unlink (new
  `s-contacts-36`); «differ» is said against what is shown; the combined
  list's panel shows the linked record (`s-contacts-1`); Анна is linked in
  three books everywhere, «Which copy to edit?» is hers (`s-contacts-32`,
  `-35`).

- **The import wizard's «Into» is a drop of links** (Carrel's task 225;
  accepted by the owner 23.10): `s-io-3`…`-8`, `-15`…`-19` — a `Drop` with
  a `Pick` trigger, the collections as links to the same step, a read-only
  one shown off with its reason, «New ‹kind›» under the line; the menus
  `s-io-10`…`-13` the same. Closed, it looks as before (pixel for pixel).
  `s-io-6` has the row its «1 with errors» counts. The notices above
  Connections' cards (`s-profile-31`, `-32`) stand with the page's gutter,
  their diagnostics too (Carrel's task 224).

- **Search: the day of an event, the states of a search, the menu of
  suggestions** (Carrel's task 189; accepted by the owner 23.10). `s-find-1`:
  an event shows «10 Aug» and «15:00 · Zoom»; a task's tick is a sign; a
  file has the collection's mark instead of a tick; names as the card holds
  them; the bar has no time or «Search again» of its own — the header's
  button searches again; the subtitle says only what was found. `s-find-2`:
  «7 found», «no source answered». `s-find-3`: the jumps carry their
  collections' marks, «Timeline · Анна **Ковал**ёва», the history keeps
  only what holds the typed. New frames: `s-find-4` (running, Stop),
  `-5` (nothing found), `-6` (the search page without a query — its own
  field), `-7` (nothing ticked), `-8` (too many searches), `-9` (no source
  answered), `-10` (stopped). The phone's search button is a link to
  `/app/search` (`phone-1`, `-4`); the mailbox's search bar and subtitle
  follow (`s-mailbox-14`).

- **Import after note 111** (Carrel's task 222; accepted by the owner 22.10,
  Р-76). A card without FN gets it from N only, else «Unnamed contact» —
  «Ready — name from N» (`s-io-5`); an event's or a task's default title is
  shown, not written. The marks of «Ready — …» are joined by commas: «a
  whole day (no end), no title» (`s-io-6`). An empty cell «—» is muted in
  every import table (`s-io-5`…`-8`, `-19`). «With a new UID» holds only
  what the answer did not do: «the file holds this UID twice», «the server
  gave its own» (`s-io-9`). «Try again» in the unanswered collection's
  notice (`s-io-19`). Words in the captions: the drop refusals by case
  (`s-io-1`), the refusals of step 1 (`s-io-18`), the report's rows
  (`s-io-9`), a set's subtitle and a zip in a set (`s-io-4`).

- **Collection sheets: addresses, names, states** (Carrel's task 192;
  accepted by the owner 22.10, Р-74, Р-75). The sheets post to C-5's
  `/app/collections/{new,rename,move,delete}?a=…&c=…` (a new folder to
  `new?kind=file&a=…`) with `account`, `components` (`VEVENT`, `VTODO`,
  `VJOURNAL`) and `name`; the address in «Rename and colour» is read-only
  and not sent. `s-profile-6`: radicale · vps refuses, the typed form rides
  hidden, «Back to the form» posts `?op=back`; `s-profile-8`, `-9`: Личный,
  with «nextcloud may tell the attendees of its meetings that they are
  cancelled.»; `s-profile-21`: no fields, «Try again». New frames:
  `s-profile-28` (no server holds calendars — Connect account), `-29` (New
  collection without a kind, «What it is»), `-30` (rename refused in the
  sheet), `-31` (delete refused, a notice on Connections), `-32` (a move
  that kept the original), `-33` (a feed's name and colour), `-34` (Move
  folder), `-35` (moving a whole storage); the whole storage's menu has Move
  (`s-profile-17`), the folder's menu items are links (`s-profile-12`).
  Operations on a collection stay in Connections (Р-75); the menu in the
  head of an open collection is deferred.

- **Import: step 1 of every kind, one rule of refusal, the report's places**
  (Carrel's task 188; accepted by the owner 22.10, Р-72, Р-73). New frames:
  `s-io-15`, `-16` — step 1 of the calendar and of the list of tasks with
  their own words; `s-io-17` — an empty «Into» from a combined view, Preview
  file off with its reason; `s-io-18` — a refusal of step 1 (the other
  refusals' words in its caption); `s-io-19` — the collection did not
  answer on the preview, Import off. The file field of every step carries
  its extensions under it; notes take `.ics` (journal entries) and several
  files (`s-io-4`); the tasks' and the calendar's previews no longer refuse
  a record without a title or an end — a default title and a short mark
  (`s-io-6`, `-7`, `-14`); one phrase for one failure in the preview and
  the report; the report has «With a new UID» and «Not kept by the server»
  (`s-io-9`); «Into» and the preview's head carry the collection's mark
  (`s-io-5`…`-9`, `-14`); the read-only notebook is in the notes' «Into»
  (`s-io-11`); the import's messages stand flush with the column. A
  contact's subtitle naming its address book carries the book's mark
  (`s-contacts-1`, `-6`, `-16`, `s-publish-3`, `mentions-1`, `-9`).

- **The letter and the links after 194** (Carrel's task 221; the cross
  accepted by the owner 22.10). «Send it as» is `format` (`md` | `ics`), the
  link letter's line `line`, the attach pickers' search `q` in a GET form
  `attach-search`; Published links' live rows: Send a link
  `/app/mail/link?id=…`, Revoke a submit of `post-revoke` (`s-publish-1`,
  `-15`, `-17`, `-20`, `s-mail-4`); the letter's line of recipients has its
  Enter; the crosses of the letter's sheets are links back. A cross that is
  a link is muted as a button cross is (`s-publish-13`, `-22`,
  `components-41` — it was ink).

- **Restore from this computer: step 2 in place, records at Restore**
  (Carrel's task 220; accepted by the owner 22.10, Р-71). `s-backup-15`,
  `-16`: the page's column is the zone «restore»; new `s-backup-20` — step 2
  of an archive from this computer: «From this computer» and «compared when
  you restore» in «What will happen»; its form asks for the parts of the
  chosen collections at Restore.

- **A folder's upload queue waits in its frames** (Carrel's task 219):
  `s-files-1` and the empty folder `states-32` carry the hidden queue of
  the upload (`data-tc-progress="post-upload"`, the row and the taken-name
  templates — Keep both · Replace · Skip) and the `post-upload` form;
  `s-files-2`'s queue is the same with the server's rows; the Files and
  Folder fields (`s-files-8`, `states-32`) ask about taken names first
  (`data-tc-taken`). Pixel for pixel the same.

- **A new device's addresses: «ro» at every read-only collection**
  (Carrel's task 217): `s-plan-4`, `-5` — «Общие с семьёй» and «Команда»
  carry «ro» beside their protocol, as Public does and as every other frame
  draws them.

- **Restore from this computer: the archive last** (Carrel's task 218;
  accepted by the owner 21.10). `s-backup-15`, `-16`: Archive is the
  card's last field, after the storage and the folder — without a script
  the parts go in the form's order, and the server knows where a large
  archive goes before it comes; «None — up to 16 MB» switches the folder
  off.

- **Publishing and mail without scripts** (Carrel's task 194; accepted by
  the owner 20–21.10). The publish sheets are forms (`how`, `from`, `to`,
  the ticks of what the visitor sees, `download`, `attachments`, `lives`);
  «Time» is a check that is off — it cannot be taken away; «Link lives
  for» is a drop of radios; «The whole calendar» switches the days off.
  The screen under a sheet is the page it is opened from: Connections
  under a calendar's, the card under a contact's, an event's, the note's
  page under a note's, Published links under «Send this link». A letter's
  recipients are tags with a cross and a line «Name or address»; Attach is
  a menu of submits; the attach sheets are frames of their own at
  `/app/mail/attach` (note, file, calendar, and new contact and folder),
  900 px, one tree on the left, choices by address; the calendar's without
  Show. «Send this link» opens from a link row's Send and from «Send by
  letter» of a link just made: a link goes in a letter's text, never among
  its attachments. New frames: a dead link's row (`s-publish-21`), a link
  with no expiry and no mail (`s-publish-22`), a note's page whose files
  were not published (`s-publish-23`). On a narrow sheet the line at the
  main button goes above the buttons (the deferred mail client's reply
  too); a picker's rail and list stack and stretch on a phone.

- **Appearance: the time zone as a line of text** (Carrel's task 193;
  accepted by the owner 19.10): «Time zone · Europe/Moscow · this
  instance» instead of a read-only field that went with the form; the
  hint under the dates and times is two short sentences.
- **Connections: «Update collection list», a row's tick that sends
  itself, the shown password on the page** (Carrel's task 193; accepted by
  the owner 19.10). Each card that answers has «Update collection list»
  after Test (the server asked again: new, renamed and gone collections;
  the name the owner's). A collection row's tick is a submit button of
  the page's `conn-picks` form (`pick` = account|collection): in the
  combined views or out of them, without a script too. `s-profile-4` is
  the whole page with nextcloud's password shown under its address, the
  unlock badge and Lock now in the subtitle; `s-profile-27` — after
  updating the list: «nextcloud: 1 new collection — «Поездки». Tick it to
  see it in the rail.» and the new row, unticked, marked `new`. Server
  URL's hint has one wording. `Badge` is no part of `Acts` any more.
- **The restore's match: a closed choice of the target in every row, and
  beside the bar's Restore into** (Carrel's task 212 pp. 6, 7; accepted by
  the owner 19.10): `s-backup-8` — `into.N` per row and `into` in the bar,
  radios of the collections with their marks, New collection, Skip this
  one; no open menu, no text with an arrow.
- **A davloom device's row** (Carrel's task 215): its access options send
  the page's `davloom-acts` form to their own address; «… collections ·
  Change» is a link to the device's scope sheet. The look is the same
  (pixel pair of `s-davloom-1`).
- **The account: the own session's End off with its reason; the password
  change in two cases** (Carrel's task 193; accepted by the owner 19.10).
  The own session's End is off — «This is the session you are using: sign
  out instead.» (`account-1`, `-15`, `s-mail-9`). `account-3` is the
  ordinary change, Back to settings a link; `account-17` the forced one:
  the red line, no way back, no rail, search or New note — only the
  user menu to sign out. Password fields are password fields.
- **The queue's answers to a taken name bring the file again** (Carrel's
  task 214 p. 1): the server keeps no file while it asks, so Keep both and
  Replace are file picks of the answer's form (`keep`, `replace`), sent at
  once with the script and by Send without one; Skip is that form's
  button (`s-files-2`, `components-26`). The look with a script is the
  same.
- **Restore from a server and its report, a new job, the escrow after a
  recovery** (Carrel's task 193 p. 5; accepted by the owner 19.10).
  `s-backup-17` — «From a server»: a collection and the archive's path in
  it with Browse; `s-backup-18` — step «3 · Report»: a row a collection,
  where it went, how many created, what happened (a new UID, a refusal in
  red), «Nothing on the server was deleted or overwritten.», Done;
  `s-backup-19` — «New job»: the job's form empty, Create, no Delete job;
  `account-16` — «Recovered · when · who: what» under «Deposited», a row a
  recovery.
- **Restore from this computer: no copy of the archive on Carrel** (Carrel's
  task 213 p. 1; accepted by the owner 19.10). `s-backup-15` — the step
  with «From this computer»: Choose a file, «Storage for a large archive»
  («None — up to 16 MB» first), «Folder in it» with «A larger archive is
  put here, read from there and deleted after the restore.»;
  `s-backup-16` — refused, larger than 16 MB with no folder, the folder
  off while the storage is «None». The archive password's note is a
  field hint now: «The one set in the job, not your sign-in password.»
  (`s-backup-7` too).
- **davloom: no devices, a device's scope, off; the tree behind a
  button** (Carrel's task 193 p. 5; accepted by the owner 19.10).
  `s-davloom-6` — «No devices yet.», the rail without a count;
  `s-davloom-7` — «… collections · Change» opens «What it sees» over
  davloom (`/app/settings/davloom/scope`), the collections ticked by
  connection, Save; `s-davloom-8` — «davloom [off]», Add device off with
  its reason, the devices kept and revocable. The tree «What the client
  sees» left the page for a sheet (`s-davloom-9`,
  `/app/settings/davloom/tree`) behind a button by Copy; a device row's
  «… collections · Change» stands at the right edge before its access.
- **The external trigger's fragments as the sheet draws them; the job
  list says what starts a job** (Carrel's task 212; accepted by the owner
  19.10). The tabs' cards (`s-backup-3`, `-4`, `-10`, `-11`) are the
  sheet's: a rubric a tab with `local`, «visible only now» only at the
  password, the tabs links to the sheet. `s-backup-5` is «What it opens»
  without the password: «The password downloads the archive too: with
  Encrypt the archive off, that is your data in the clear.». «What to
  copy» (`s-backup-2`) has baikal; the archive password's alarm is
  «Separate from the sign-in password.». The job list (`s-backup-1` and
  under the sheets) tells the home job as the form does — four
  collections and files — and «When» says only what is set: the sign-in
  rule, a `trigger` badge when an external password is issued, «—»
  otherwise (no «cron · 03:00», no «on demand»).
- **The upload sheet without «Open the folder»; the connection dialogs
  and the new device from the set** (Carrel's tasks 210, 211; accepted by
  the owner 19.10). `s-files-19`: its tree's rows are links, no button.
  The connect dialog (`s-profile-2`, `-11`, `-24`) carries «The server asks
  no password», unticked, a password field, the server's field names.
  Editing a connection (`s-profile-16`, `-22`, `-23`) edits nextcloud: its
  label, «Leave it empty to keep the stored one.» under the password. The
  feed's label colour (`s-profile-10`) is a choice of six. The new device's
  tabs (`s-plan-1`, `-2`, `-4`, `-5`, `-6`): an account a row with its
  count, baikal off with «password refused»; Thunderbird and Anything else
  a row a collection, Anything else with the WebDAV ones; the Android code
  a picture from the server.
- **Attachments and a new device: the states without frames** (Carrel's
  task 193 p. 5; accepted by the owner 19.10): `attachments-5` — no folder
  yet, «Folder [not set]», the collection to choose, the path a
  placeholder, «Larger — wait for the folder below», Open in Files off;
  `attachments-6` — the folder chosen but refused, «In this folder now
  [unavailable]», the path and the server's answer instead of the counts,
  Open in Files off with the reason; `s-plan-6` — a new device while the
  unlock window is closed: no countdown, «Asks for your Carrel password
  first.», no link card yet.
- **Connections: the states without frames** (Carrel's task 193 p. 5;
  accepted by the owner 19.10): `s-profile-24` — the connect dialog after
  the server refused the login: «Could not connect. … Nothing was
  saved.», the discovery trace, what was typed kept, the password not
  given back; `s-profile-25` — nothing connected: «Nothing is connected
  yet.», the rail without counts and statuses; `s-profile-26` — the
  Disconnect sheet at `/app/settings/connections/disconnect`: nothing is
  deleted on the server, what Carrel removes in numbers without names.
  `s-profile-2` takes the server's field names, a password field and
  Cancel as a link.
- **The external trigger is a sheet over the job list, with its password
  as a row** (Carrel's task 207; accepted by the owner 19.10): «External
  trigger» in a job's «⋯» menu opens `/app/settings/backup/trigger` — off
  (`s-backup-14`: the price, Turn on), on (`s-backup-12`: the password
  row — issued, last run and from where, runs — with New password and
  Revoke; the tabs and the line with the password hidden), issued
  (`s-backup-13`: the password and the line whole, «visible only now»).
  The commands are seen without opening the job's form; the job's page
  (`s-backup-2`) has no trigger card. The history says «external», the
  password and the address of each run, a refused one in red
  (`s-backup-1`). The script of a job that also downloads (`s-backup-3`)
  and the GoSentry, Gitea and curl tabs are fragments of the sheet; Run
  now is gone from the job's head.
- **The Email block says its state with a badge in its title** (accepted
  by the owner 19.10): «not confirmed» (alert), «not confirmed» and «link
  sent 12 February 11:42», «confirmed» (no date), «pending» and
  «new@example.org» or «new@example.org · link not sent» — the words
  beside the badge plain, since a title prints its badge in capitals. No
  red box under the address: Send confirmation and Send again stand
  after Change, the field 420 px for the two buttons (`s-mail-9`,
  `account-1`, `-8`, `-9`, `-10`, `-15`).
- **Settings against themselves** (Carrel's task 209): `s-mail-9`
  draws the account page whole — Password, Sessions, Key escrow and
  Delete account under its Email block, as `account-1`. The connections
  (`s-profile-1`, `-10`) draw the set the other frames draw: nextcloud
  with nine collections (Личный, Дела, Работа, Дни рождения, Записи,
  Дневник, Общие с семьёй, Контакты, Архив) and Files, radicale · vps
  with four (Проекты, Дежурства, Команда, Рассылки), baikal, then
  sftpgo · vps; the numbers of the fixtures; «Праздники РФ» lands in
  Calendar and Tasks (3 and 1), the league's cache holds 2 events.
- **Administration against its data** (Carrel's task 208; accepted by
  the owner 19.10): letters from users only through configured mail —
  in `s-admin-3` (mail not configured) the check is off with «Mail is not
  configured: there is nothing to send through yet.», and with sending
  off there is no Sender address and no limits (`s-admin-3`, `-33`);
  `s-admin-32` has sending on, its sender the From address
  `carrel@slovo.example` (no made-up `noreply@carrel.example`). The DAV
  check's second stage (`s-admin-29`): the file row skipped, grey as an
  unmeasured row of the installation check — «skipped: no collection of
  this kind that can be written to» (the trace says 0 plain folders).
- **Files: what cannot be done is off, and one tree for one set** (Carrel's
  task 206; accepted by the owner 19.10): in a read-only storage
  (`states-34`) the rail's New folder is off, its reason «This storage is
  read-only.»; the subtitle there and in `s-files-15` says it with a
  `read-only` badge instead of a sentence. The upload sheet's tree
  (`s-files-19`) has the read-only Public row, dimmed, as `s-notes-11`
  has. The sheet «Attach from Files» (`s-files-16`) is `picker-sheet`,
  and its servers' menu swaps it at `/app/files/picker?for=attach`;
  the upload sheet of `s-notes-11` is `upload-sheet`, the id its
  links swap.
- **The password sheet, one for every handing out** (Carrel's task 193
  p. 3, 4; accepted by the owner 19.10): its title says what for, the
  subject whose, the main button the act; a password field; one line
  «Unlocks stored credentials for five minutes.»; Cancel a link to the
  screen under it. `s-plan-3` is «Generate profile» (it said «Show
  password · nextcloud»). New `s-davloom-5` (over the device form, what
  was typed kept, «Create password») and `account-15` (Delete account: the
  password is the confirmation; no line about five minutes). «Carrel
  never touches it» is gone from the account's deletion: not true.
- **The backup job's form** (`s-backup-2`, `-3`; Carrel's task 193 p. 1,
  2; accepted by the owner 19.10): «What to copy» — a column a connection,
  collections are checkboxes `src` of the form, a connection's tick only
  shows how many; the destination a pick with the collections' marks; the
  folder a field with Browse; one line about the archive's forms;
  «When» — one checkbox (on sign-in) and a list of the other ways to start
  it; the trigger card is the job page's, not the form's; no Run now in
  the form's head; the archive password a password field.
- **Administration without polling, the cards whole** (Carrel's task 195
  p. 8; accepted by the owner 18.10): no «2 min ago» in any administration
  frame (19); Users has six rows for «6 accounts · 2 administrators · 4
  covered» and no line under the table; `s-admin-18` keeps Role and Create
  link with only Send invitation off; `s-admin-19` keeps Role and the
  warning, its fields on one line; the optional email is empty, no «—»;
  the forms name their fields as the server does.
- The arrow of a pick (and of a folded switch) stands 1.5 px higher: its
  middle on the middle of the words beside it (the owner, 18.10). No box
  moves.
- **Files against «what cannot be done is off»** (Carrel's task 205):
  New folder off in the root's rail (`states-35`), Details off until one
  file is chosen (`s-files-2`, `states-32`…`-34`), Rename, Move and Delete
  off in a read-only storage's selection bar (`states-34`); the head's
  tick `is-part` when a part is chosen (`s-files-1`); Clear selection with
  tiles (`s-files-5`); a refused upload in the page's queue (`s-files-2`);
  an empty folder's Upload picks files itself, Send without a script
  (`states-32`). Sheets: the connection in the head is a drop of links
  (new `s-files-26` draws it open); the copy sheet is the move sheet's
  form with Copy and «Open the folder» (`s-files-11`); the New folder
  refusal keeps the sheet's message (`s-files-13`); `s-files-12`'s message
  as in the sheet; the upload picker's list of picked files is empty as
  the server prints it, no cross, its tree links to the sheet over that
  folder (`s-notes-11`, `s-files-19`); no line under the tree — the
  attachment folder's badge says it; the refusal in the upload sheet 8 px
  from the field (`MsgFlushY8`). New `s-publish-20`: the published links'
  page without mail. Accepted by the owner 18.10. On a phone a selection
  bar wraps what does not fit (Р-65): Clear selection takes a row of its
  own.
- New `s-admin-34`: deleting an account — a sheet with its own address
  over Users, what goes said before the button, the login typed to
  confirm; Delete in the row's «⋯» (`s-admin-12`) leads there (Carrel's
  task 195 p. 6; accepted by the owner 18.10).
- `s-admin-3`: the settings form names its fields as the server does;
  «Send test» is never off — it sends what is typed, and with no Host its
  answer says so; a password field; a wider test address. Extra frames
  `s-admin-32` (mail configured, the relay's answer to the test) and `-33`
  (a field refused: «A port from 1 to 65535.» — the bounds of a number are
  said by its refusal; the test not sent) — Carrel's task 195 p. 5;
  accepted by the owner 18.10.
- **Key escrow and the password's history** (Carrel's tasks 195 p. 4, 202
  p. 1; accepted by the owner 18.10). `s-admin-6`: the forms are one page
  form each (`action` hidden), password fields, «… logged and shown to
  them», the line under Recover names the history link, six accounts for
  «4 of 6». Extra frames `s-admin-30` (recovered: the temporary password
  once, and whether a letter went) and `s-admin-31` (escrow off: the
  master password twice, «Turn on key escrow»). `account-4`: the
  permanent link to the password's history with Copy; shorter words by
  Withdraw. New `s-auth-15`: the history page by that link, without
  sign-in.
- `s-admin-5`: the DAV validator's fields are its form (server, login, a
  password field — «Asked each time, never kept.» —, «Certificate
  fingerprint · SHA-256 · optional» as in a connection, Carrel's task
  201); Test discovery, its trace, then the second stage: the consent and
  «Run the full check». Extra frames `s-admin-27` (before a check),
  `-28` (a broken step), `-29` (the second stage's report, a row a kind)
  — Carrel's task 195 p. 8; accepted by the owner 18.10.
- `s-notes-2`, `-20` say in their labels that they draw the full width a
  person chose in Appearance; the default is the reading width
  (`s-notes-3`, `appearance-1`) — Carrel's task 204, the owner 18.10.
- `s-admin-4`: the address and «Run the check» are one form (the field and
  the button name the page's `tc-post` form `install`); «1 not measured»
  instead of «1 warning»; the rail says «2 checks failed». Extra frames
  `s-admin-25` (never run: no table, «Not checked yet» in the rail) and
  `s-admin-26` (an address that is no address: the field in error, the
  reason under the head, no table). No frame for «running»: the check is
  one request (Carrel's task 195 p. 1; accepted by the owner 18.10).
- `s-admin-7`: the journal's bar is a form — the action and the day
  («since», from that day on) go together, Show without a script; the
  segments are links. A login that succeeded has no address; refusals in a
  row by one account are one line with a count, «3rd attempt · ip
  203.0.113.9 · rate-limited» (the owner, 16.10). `s-admin-10`: the
  actions are radios of that form (Carrel's task 195 p. 2, 3; accepted by
  the owner 17.10).
- **The conflict screen's states** (Carrel's tasks 191, 198; the owner's
  decisions Р-67, Р-69; accepted by the owner 17.10). The form's names follow C-5: radios `choice.0…` with
  `mine` / `theirs` («Leave it out» is `theirs`), hidden `etag` and
  `op=combine`; «Use only my edit» and «Use only the server version» submit
  forms of their own (`conflict.mine`, `conflict.theirs`). Extra frames:
  `s-conflict-6` — whose change is not known: «Your form» and «Use only my
  form», no kind of difference, nothing chosen, «Choose a side for every
  property.» under the title, «Write this combination» waits for a side of
  every card; `-7` — deleted on the server while edited: the edit («write
  it back») against the deleted record («leave it deleted»), nothing
  chosen, «Write this choice» waits; `-8` — a combination that cannot be written, its
  fields; `-9` — the server refused the write; `-10` — the server did not
  answer; `-11` — a note's move that waits. When nothing differs there is
  no conflict screen: nothing is asked (the owner, 17.10).
- The collection's radio items in an edit's head name the edit's form
  (`form="…"`, Carrel's task 203): `s-tasks-3`, `-14`, `s-agenda-3`, `-15`,
  `-27`, `s-contacts-3`, `-4`, `-27`, `-31`, `states-12`, `-17`, `-24`,
  `-36` (the notes and `s-contacts-15` named it already). The look is the
  same.
- `s-profile-5`: the address of a new collection shows the derived
  `poezdki` as its placeholder, and the path under it ends with it (the
  owner's decision Кл5, Carrel's task 200).
- **The series question and the rewrite sheet over their screens**
  (Carrel's tasks 184, 185, 197 p. 2; accepted by the owner 17.10).
  `s-agenda-27` (new, extra): the question before saving a series is the
  answer to the edit's POST — over the edit with what was typed (until
  18 January 2027); «What you changed» says what changed, shortly.
  `s-agenda-28` (new, extra): the series written in part, over its card;
  «Try again» goes with `op=apart` and the days in its address
  (`s-agenda-24` too). `s-contacts-15`: the rewrite sheet is one sheet for
  any number of replacements — lines «old → choice», the groups of risk as
  ticks, all ticked by default («Rewrite 20 records»), a form;
  `s-contacts-29` is an example of an answer (one group unticked).
- **Sheets over a screen with their own address; second buttons without a
  script** (Carrel's tasks 183, 187, 197, 199; accepted by the owner
  16.10). New route frames: «Which copy to edit?» over the address book
  (`s-contacts-35`), «Which copy to mark done?» over the task list
  (`s-tasks-19`), the password reset over Users (`s-admin-24`); the quick
  sheets of an event and a task, and the escrow deposit, carry their own
  address; «nothing to write into» (`s-notes-6`) is a state, not a route.
  Under a sheet is the screen it was opened from, the same body: the move
  sheet's rail foot is the book's (`s-contacts-2`), its books go in the
  rail's order, the account named in full. Second buttons of forms have an
  address of their own — `formaction` (Test only, Check the address, Open
  the card, Leave as is, Attach, Preview, Copy link instead, Send
  invitation, Reset remembered view settings), a name and value (Cancel of
  the series question), a link (Cancel of every sheet, Leave it in the
  series, Download .ics first, Open in Files, Done); Copy buttons say what
  they copy, their fields are read-only. The import preview's rail shows
  the unticked calendar too (`s-io-6`); the refusal to create a collection
  no longer promises what Carrel cannot tell (`s-profile-6`).
- **Where a record is saved, and its collection's mark** (the owner's
  decisions of 16.10; Carrel's task 182). A record's head carries the
  collection's mark beside the collection's name under the title, not at
  the title (16 heads of cards and panels; a collection's own head keeps
  its mark at its name). An edit does not name its collection under the
  title: the choice of it is in the head, left of Cancel — «List»,
  «Calendar», and now «Address book» for a contact too (moving is common,
  Р-63); a note's «Notebook» stands right of Cancel in its bar. 13 edit
  frames of tasks, events and contacts, 5 of notes; the frames of states
  caught up with their forms (List, Attendees, Link, Repeats of a task;
  Calendar of an event). The line «Saving changes this copy only…» is
  gone from the form of a chosen copy (`s-contacts-31`) — the field says
  it. A task's copy differs in «Due», not «Time» (`s-tasks-2`).
- **The edit forms of a task, an event, a note and a contact** (Carrel's
  tasks 171–173, 175, 179–181, accepted by the owner 16.10). A task's edit
  (`s-tasks-3`, new `s-tasks-14`): «List» beside the summary — saving into
  another list moves it; «Attendees», the event's field (Р-64); «Link» in
  «Other»; «Repeats», the event's rubric, its hint saying a task repeats
  from its start date or its due date; the new task's «When» has «Time
  zone» (a task all day has none, as an event). An event's edit
  (`s-agenda-3`, `-15`): «Calendar» last in the head, «Link» in «Other».
  A note's fields (`s-notes-4`, `-20`, `-26`, `-30`, `states-31`):
  «Link» after Tags. The task columns menu (`s-tasks-12`): «Attendees»
  last, off. A contact: «Websites» rows as the email's, after «Addresses»;
  «Dates» rows — what a date marks and the date — instead of one
  «Anniversary» (`s-contacts-34`, new): vCard has one ANNIVERSARY and no
  label, so the server writes them as Apple does (`X-ABDATE` with
  `X-ABLabel`). Empty, both fold into «Show N empty fields» of a card
  (`s-contacts-3`, `-4`, `-15`, `states-12`: 7), a new contact shows them
  (`s-contacts-27`).
- **Publishing, a connection's links, a letter's links** (Carrel's tasks
  163–165, accepted by the owner 16.10). `s-publish-2`: «What the visitor
  sees» — time always, title on, place, description, attendees,
  attachments and tags to tick; without the title an event shows as
  «Busy»; «Show event descriptions» goes into it. `s-publish-19` (new):
  the same sheet for one event or task. `s-profile-22` (new, an extra
  frame): the address of a connection with published links is being
  edited — «Published links», keep or revoke, comes at once
  (`Field.Changed`); `s-profile-23` (new, an extra frame): Save without
  an answer — the dialog comes back with the reason in red. A label,
  password or fingerprint moves the links silently. `s-mail-1`: a file
  that goes as a link shows its term («30 days») and «Change…», which
  opens the publish sheet over the letter.
- **Deleting a record that has something to lose; the edit forms' tails**
  (Carrel's tasks 166 and 169, accepted by the owner 16.10).
  `s-agenda-26` (new): the sheet of deleting an event in full — no copy
  in Carrel, what points at it (a published link, related records), that
  the server may tell the attendees, which attachments go and which stay
  (a linked file may be ticked to go too), its copies in other calendars
  by ticks; «Delete» in the danger tone. `s-tasks-18` (new): the simple
  case, one line; a record with nothing to lose is deleted at once, as
  before. A task's edit (`s-tasks-3`) gets «Kept as the server sent them»,
  an event's (`s-agenda-3`) «Source», so both forms end alike. The
  event's path is Nextcloud's (`/remote.php/dav/…`) in its card too.
- **The buttons of a card, after Carrel's tasks 167–174** (accepted by the
  owner 15.10). The cards of an event, a task and a note (`s-agenda-2`,
  `s-tasks-2`, `s-notes-2`) get «Send», as a contact's has; «New task» goes
  from the cards (`s-contacts-6`, `s-agenda-2`, `s-tasks-2`, `s-notes-2`) —
  a task is made in tasks. «Print» is of the third rank in a card's head:
  on a phone it goes into «⋯», its first line (`s-contacts-6`, `-8`, `-10`,
  `s-agenda-4`, `-5`, `s-tasks-4`, `-6`, `s-notes-14`). A note's «Export
  Markdown» is the export icon and «.md», «Source» the code icon alone, the
  width switch of the third rank (`s-notes-2`, `-3`, `-27`); a note's «⋯»
  exports «.ics» too (`s-notes-14`, task 174).
- **Icons that read at once** (accepted 15.10): «Attach» has the paperclip
  in every frame, 31 buttons; Import and Export the new `i-import` and
  `i-export` (57 and 37), Upload `i-upload` (39), the rest of Download
  `i-download` (14).
- **Reliability on a phone** (Р-65, Р-66; the cards, the run bar and the icons accepted by the owner
  15.10): the date fields of the timeline (`s-contacts-7`, `s-mailbox-21`)
  and of the audit log (`s-admin-7`) hold their calendar; the bars of
  tasks, files, the audit log and the others wrap instead of being cut;
  the jobs of `s-backup-1` are cards on a phone with «Last backup» under
  the name (`Run`, `Edit` in «⋯»), and its history says «read back · 1 not
  copied» with the reason in the hint, Restore stays; the search of
  `s-find-1` keeps its query in a row of its own, its run bar says «13 of
  14 · 1 failed · 1.8 s · 11:42 · ↻» on a phone, and its line about files
  is one sentence for a person. New icons: `i-download` (an arrow into a
  tray) on every Download button, 32; `i-restore` (an arrow back round) on
  every Restore button, 7.
- **Files after Carrel's second scan** (task 148). `s-files-5`: the tiles go
  as the table does — folders first, then files by name. `s-files-16`: the
  folder «attachments» holds what every other frame draws in it — the folder
  `2026-07` and four files by name. `s-publish-18` (new): the list of
  published links when there are none — one line under the head that says so
  and where a link is made; no table, no button, no count in the rail. The
  caption of `s-publish-3` says «Photo» decides about the photo in the `.vcf`
  only. `components-40`: the thumbnail keeps its icon under the picture; the
  frame looks the same. 429 frames.
- **A folder's tile is not a file's** (the owner's decision of 12.10): no
  card under the icon, a large folder in the accent colour —
  `.tc-thumb--folder`, `Thumb{Form: ThumbFolder}`; a folder with
  `ThumbOnCard` still prints, as a file does. `s-files-5`, `dark-3` (its
  folder stands first too).
- **Lists of tasks and notes after Carrel's tasks 167 and 170** (the owner
  accepted 15.10; Р-63). Tasks get «Filter by word» first in the bar, as
  notes have — a task has no sort by title: the date matters there (the
  server's «Title» order is not drawn). Notes get a fourth order, «Recently
  changed»; their orders hide under 800 px of the column, as the tasks'
  do (`is-2nd`) — four did not fit at 460. All the bars of both kinds:
  `s-tasks-1`, `-13`, `-15`, `states-19`, `-20`, `-23`, `s-notes-1`, `-31`,
  `states-26`, `-27`, `-30`, `phone-8`. The quick sheet of a task
  (`s-notes-10`) has «Remind» after «Due time», the event's list, «No
  reminder» by default, and the event sheet's width. The columns menu of
  tasks (`s-tasks-12`) has «Where» last, off.
- **A field of a bar is as tall as the bar on a narrow screen**: in the
  drawer mode a field kept its size class (`tc-in--h26-…`) and stood 26 px
  among 38 px buttons. `.tc-bar .tc-in.tc-in` and the same in a docbar and
  a head's actions are 38 px. Lists with a filter at 641 px and less:
  `s-contacts-1`, `-2`, `s-files-1`, `-2`, `-4`, `-5`, `s-io-1`, `s-admin-2`,
  `-7`, `dark-2`, `-3`, `states-4`, `-8`, `-9`, `-34`.
- **The Save bar in every frame of a loss report** (Carrel's task 162; the
  kinds are equal): `states-12` (a contact), `states-17` and `states-36` (an
  event) carry the edit's hidden Save bar under the head, as `states-24`
  and `states-31` do — the edit after the report is the same form with the
  same Save. Pixels the same at 1351, 820 and 460.
- **The other cards of a linked record from the combined view** (after
  `s-contacts-33`, Р-58; accepted by the owner 15.10). An event
  (`s-agenda-25`: the place taken from «Личный», its mark after it), a task
  (`s-tasks-17`: the due date from «Дела», the priority from «Работа») and
  a note (`s-notes-32`: «text from Записи» in the subtitle, the tags from
  «Дневник») — new fragments. The mark of «Дневник» in the blocks of copies
  is violet, `#6B4E7A`, as in the rail (it was green): `s-notes-2`, `-3`,
  `-27`. `DocRefs` prints no empty note after a copy that holds the same.
  446 frames.
- **Two cards of a linked record** (Carrel's task 153; the owner's
  decisions of 13.10, Р-58). Opened from the combined view, the card shows
  what the person chose in the duplicates section, and a value where there
  was a choice carries, after it, the mark and the name of the collection
  it was taken from — `s-contacts-33`, new. Opened from a collection, the
  card shows that collection's own data, as its edit does, and a field
  where there could be a choice carries the badge «linked» —
  `s-contacts-6`, `s-agenda-2`, `s-tasks-1`, `-2`, `s-notes-2`, `-3`, `-27`
  (a note's badge stands at the rubric «Tags»). A record is linked only by
  such a choice; without one it is a candidate. `FieldsItem.After` holds
  what stands after a value that is a link, outside the link; a `Mark` is a
  part of a value. The block «The same note in other notebooks» is gone
  from a note's edit (`s-notes-20`, `-26`, `states-31`) and stays in its
  reading. 442 frames.
- **Which copy of a linked record is edited or marked done** (Carrel's task
  153; the owner's decisions of 11.10 and 13.10, Р-58). An edit of a record
  with linked copies is written into one copy, and the person chooses which
  before the form: Edit opens the sheet «Which copy to edit?» —
  `s-contacts-32`, new. The Done tick of a linked task in a list opens
  «Which copy to mark done?» — `s-tasks-16`, new — for the task's row
  stands for all its copies. In both sheets a copy is a tick of the sheet's
  form (a radio button for the edit, a checkbox for Done) beside its mark
  and what differs, and the main button acts — «Edit this copy», «Mark
  done»; a read-only copy is a row without a tick, with its reason.
  `s-contacts-31`, new: the form of the chosen copy says where the edit
  goes, «Saving changes this copy only — …», in one line under the head. A
  record without linked copies has no sheet. 441 frames.
- **The sheet that asks before a save** (Carrel's task 151; the owner's
  decision of 11.10, accepted 13.10). Where an edit does not show what the
  person meant and a wrong guess loses their data, one sheet asks before
  anything is written. `s-agenda-22`: a series was cut short or its rule
  changed, and the meetings moved by hand are left without their day —
  «Keep the meetings that were moved by hand?», what was changed, the list
  with its count, and three answers: «Keep as events» (the default),
  «Remove», «Cancel». `s-agenda-23`: one meeting, the singular.
  `s-agenda-24`: a separate event could not be written — what was written,
  what was not and why; it stays in the series; «Try again» or «Leave it in
  the series». The sheet is the pattern for other such questions: a
  question, what changed, the list of what is touched, two or three
  answers, one of them the default. No new block. 438 frames.
- **«Rewrite the addresses»: the counts agree, and the button counts on the
  spot** (Р-62). s-contacts-29: «22 records name it» (it was 14, fewer than
  the groups hold), and the button «Rewrite 17 records» (it was 9): of the
  22, 2 in Shared are read-only and 3 stand in the unticked «Meetings
  organised by others». The sheet carries its tally; the caption says how
  it counts.
- **The last selections become forms** (a tail of Carrel's tasks 148, 149;
  the owner's decision of 12.10 to carry forms through every frame).
  s-backup-8: the dialog is a form — the rows' ticks are checkboxes
  `collection`, «Restore into», «Skip» and «Restore 3 collections» send it
  with `op`, «Select all» shows only with a script (`tc-js`). phone-6 and
  dark-3 are the files' form as s-files-1 and s-files-5 draw it: ticks
  `item` of `files-sel`, the bar sends it with `op`, the tiles of dark-3 have
  their box and their name as a link; «Clear selection» on the phone shows
  only with a script. s-mailbox-3 (mail is put off) and components-16, -17
  (the bar alone, no list to send) stay as they are. `tc-js` and `tc-nojs`
  are mixins of a block, as `tc-push`: a button that carries one is still a
  button.
- **The agenda's badges stand side by side** (the owner's decision of
  14.10, Р-61). «Синк-ап по релизу» repeats weekly, so its row has both
  badges, «4 attendees» and «repeats weekly», in every frame of that week
  (s-agenda-1, -17, s-notes-5, -7, -8, -11, s-find-3, s-io-2, account-5,
  -6, s-profile-5, -6, states-6, -13, dark-1, desktop-1). The badges
  column is 168 px, so «repeats weekly» is no longer cut at 88 px and the
  title column is 48 px narrower in every agenda list (also s-dup-1,
  s-find-1, s-publish-2, -6, -13, states-3, -5). The agenda in columns
  (s-agenda-20, -21) draws badges, not words, in «Attendees and status».
- **The days of a series a server lost, in words** (Carrel's task 151, the
  question left open; the owner's answer of 14.10, Р-60). `states-36`
  (new), «Событие · сервер не сохранил дни серии»: the form of
  `states-17`, and the report after the save names the days instead of
  `EXDATE` and `RECURRENCE-ID` — «The change was saved, but the server did
  not keep two days of the series.», then each day and what became of it:
  a skip lost, the meeting is back; a move lost, the meeting is back on its
  day and hour. Properties that have no day keep their names, as in
  `states-17` and `states-24`. A shape of `tc-msg` without a mono part; no
  new block. 443 frames.
- **The agenda in chosen columns** (Carrel's task 154; the owner, 13.10).
  `s-agenda-20`, new: the agenda as rows of columns in its default set —
  Time, Title, Where, «Attendees and status», Source. `s-agenda-21`, new: an
  example of another choice, not a default — Where and Source off, Tags and
  Changed on. The row is `ColumnRow` as it is: the time is `ColumnDate`,
  first, before the name; where is `ColumnDetail`; attendees and status are
  `ColumnStatus`, in words, not a badge. The days' groups stay. 435 frames.
- **The tails of task 149 and task 155** (the owner accepted them, 13.10).
  `s-backup-7`, the first step of a restore, is a form: the dialog sends,
  «Where the archive is» is a segment of links, the table's ticks are radio
  buttons, the password has its name; it looks the same. `s-mail-9`: «Name»
  has its form and «Save», as the frame beside it. An attachment that lies
  in Files has a plain «Unlink» (hint: «Remove the link; the file stays in
  …») in place of the alarming «Remove» — `s-notes-20`, `-26`, `states-31`.
  The row of «Заметка со встречи 10.08» carries «linked ×2» where its
  neighbour does (`s-notes-1`, `-10`, `s-publish-4`): six frames draw that
  note with a linked copy, and the list now agrees with them.
- **«Rewrite the address»: several replacements, groups by a tick**
  (Carrel's task 157; the owner's decisions of 12.10, accepted 13.10).
  `s-contacts-29`, new: each replacement is the old address with its count
  and, under it, a pick of the card's addresses and «Do not rewrite»; each
  group of risk is a tick with its count as a button that opens the list of
  its records; an unticked group is left out and the button counts the
  rest. A record that stands in two groups is rewritten only when both are
  ticked, and the sheet says so. `s-contacts-30`, new: that list — the
  collection's mark, the title, the date; a row is a link. No new block:
  `Line` with `Check` and `Drop`, `Pick` with options, a menu of links. 433
  frames.
- **The steps of a wizard stand in one place** (the owner's remark of
  13.10): a step whose content is narrow is a dialog as wide as the wizard's
  widest step, and what is under the steps and the head keeps its own width
  — `Dialog{Variant: DialogStep520 | DialogStep560 | DialogStep760}`
  (`tc-dialog is-wide tc-dialog--step520` …). `s-io-3`, `-4`, `-9`, `-14`,
  `s-backup-7`: the steps and the title no longer jump by 280 px between a
  step with a table and a step without one. Use it for every step of a
  wizard that has a `DialogWide` step.
- **The import report names what was already here** (Carrel's task 159):
  «Created 117» and the card «Already here, skipped 5» with a name and a
  UID, in place of «UID collisions» with arrows (`s-io-9`, `s-io-14`); the
  card is «Replaced» or «Imported as copies» for the other answers. The
  caption of `s-io-5` gives the words of the other two answers and of
  «nothing to import».
- **Import asks about what is already here** (Carrel's task 159; the
  owner's decision of 12.10, accepted 13.10): the four previews
  (`s-io-5`…`-8`) have, beside «Into», the question «5 records of this file
  are already in Контакты» with a segment of links — «Skip them» (chosen),
  «Replace them», «Import as copies» — and a hint that says what the chosen
  one does. The count reads «already here», not «UID collisions»; «ready»
  and the button count what will be written; such a row says «Already here
  — skipped»; the step's message no longer says import always creates.
- **A reminder of one's own** (Carrel's task 156; the owner's decision of
  12.10, accepted 13.10): the lists of a reminder have «Custom…»
  (`s-agenda-9`, `s-notes-18`, `desktop-18`), and `s-agenda-19`, new, draws
  what it opens — beside the pick stand «How long» and «Unit» (minutes,
  hours, days, weeks before), and for an all-day event «At», the hour. 431
  frames.
- **An event has «Related to»** (Carrel's task 160; the owner accepted it,
  13.10): the edit and the new form of an event have the group «Other» with
  the field «Related to» before «Attachments», as a task's form has —
  `s-agenda-3` with a pill, `s-agenda-15` empty. The card and the panel of
  an event drew «Related» already (`s-agenda-2`, `s-agenda-1`).
- **The move sheet as it opens** (Carrel's task 158; the owner accepted it,
  12.10): `s-files-25`, new — the folder the things lie in is the chosen
  one, with the note «the items are here»; it can be ticked, for «New
  folder» makes a folder in it; Move is disabled with its reason beside it,
  «These items are already here.» (the owner's rule of 09.10). 430 frames.
  The note stands at that folder always — with another folder chosen and
  in the copy sheet too (the owner's decision of 31.10, Carrel's task 248):
  `s-files-4`, `-11`, `-13`.
- **A tree's folder opens without a script** (Carrel's task 158): the
  sheets whose tree chooses a folder have a second button of their form,
  «Open the folder» (`Button.NoScript`, `name="open"`), which only a page
  without a script shows — the server answers with the same sheet, the
  ticked node chosen and opened (`s-files-4`, `-13`, `-19`, `-25`). «Send
  test» of the administrator's settings is a button of the page's form with
  an address of its own (`s-admin-3`): it sends what is typed, saved or not,
  and saves nothing. The caption of `appearance-1` says where the theme, the
  density and the columns are kept: the server puts them into a cookie of
  this browser. The frames look the same.
- **A button beside a rubric's words stands level with them** (the owner's
  remark of 12.10): `.tc-rubric` that holds a button or a menu is a flex row
  centred across (`:has(> .tc-btn)`, `:has(> .tc-drop)`), gap 4 px. The
  words move up by a pixel or two where they hung below the button's middle
  («Where to put it · New folder»); a rubric with a menu («Attachments ·
  Attach») moves by less than a pixel.
- **The files' ticks and selection bar are a form** (Carrel's task 148):
  `s-files-1`, `-2`, `-4` — the ticks of the table are checkboxes named
  `item`, `s-files-5` — the tiles' boxes, all of the form `files-sel` that
  stands after the list; the bar's Download, Rename, Move, Copy and Delete
  send it with `op`. `states-34` (a read-only storage, nothing ticked) has
  the bar of an empty selection, «With the ticked» with Download and Copy,
  which only a page without a script shows. The frames look the same, to the
  pixel. The names of the fields are an illustration.
- **The settings rail counts what its pages draw** (Carrel's task 148; the
  owner left the numbers to the mockup, 12.10). The Connections page draws a
  fourth server, radicale · vps, with the collections the record frames
  already name (`s-profile-1`, `-10`); the heads say «4 servers · 15
  collections», the rails «Connections 4» (6 in the four mail frames, with
  the two mailboxes). The davloom page draws a third device, as the rail and
  the rail's foot count (`s-davloom-1`). «Backup 3» is its three jobs. 39
  frames, numbers and rows only.
- The owner's decision of 07.10 (Р-37, the other sheets): the sheets of a
  move, a copy and Attach from Files have the pick of the server in their
  head, and their tree shows the storages of that server only, without the
  marks of the server (s-files-4, -11, -13, -16); the tree of «Attachments ·
  Browse» lost the storage of the other server (attachments-4).
- The owner's decision of 07.10 (Р-37, second part): the tree of the upload
  sheet shows the storages of the server chosen in the head and their
  folders — the storages of other servers and the marks of the server on the
  rows are gone (s-notes-11, s-files-19).
- The owner's decisions of 07.10. A count and «ro» stand apart in a rail
  (Р-34): a subscription has its count — «—» when not ticked — and «ro» as a
  badge beside it, as a read-only book of contacts has; so has the read-only
  notebook (the rails of the calendar, tasks and notes, 30 rows). Dated
  folders lie in «attachments» (Р-35): the trees of the file sheets draw
  «2026-07» and «2026-08» under it, the crumbs of the tiles say
  Documents / carrel / attachments / 2026-07, a taken name is refused «in
  attachments already» (s-files-4, -5, -11, -13, -19, s-notes-11, dark-3).
  The head of the upload sheet has the pick of the server and the full path
  of the chosen folder beside it (Carrel's task 120; s-notes-11, s-files-19);
  what the pick opens is the new frame «Upload · сервер» (s-notes-29).
  397 frames.
- Carrel's task 116: the files of «attachments» in the sheet From Files
  stand a level under their folder, not beside it (s-files-16).
- Carrel's tasks 105 and 127: the rail of notes has its third notebook, the
  read-only «Общие с семьёй», not ticked — the head's «2 of 3 notebooks» and
  the Notebook pick of an edit name it (s-notes-1, -2, -3, -10, -20, -26,
  -27, s-conflict-5, s-io-4, -8, s-publish-4, phone-8, -9); the note's panel shows
  what its page shows — the second task done and struck, two related records
  (s-notes-1; `SecTasks` strikes a done task as `DocTasks` does).
- Carrel's task 128 (the owner's decision of 03.10): a folder's count and a
  picture's dimensions stand only where they are known without a request of
  their own — a folder without a count is «folder» in the list (s-files-1)
  and has an empty meta on its tile, «0 items» is a read empty one, a tile of
  a picture not chosen shows its size alone (s-files-5); the chosen picture's
  panel has Dimensions after Type (s-files-1).
- Carrel's task 113: search draws its refused source as every list does —
  the mark in the rail, the reason in its hint, and in the poll bar a red
  «1 failed» whose menu says what was asked and what the server answered,
  with Retry for that source (s-find-1).
- Carrel's tasks 112, 117: every list's «⋯» opens its menu without a script —
  the timeline (s-contacts-7: Columns; its menu frame s-contacts-14 now holds
  only what goes under «⋯», the kinds fold in their own pick), the states of
  contacts (states-4, -8, -9, s-contacts-2 — as s-contacts-1), the folder
  (s-files-1: Name / Size / Changed). The order of contacts and of a folder
  and the folder's view are links (`?sort=…`, `?view=…`), as their twins.
- A segment's option that is a link is muted as a button option is; the
  root's rule for links (`color: inherit`) had drawn it in ink (the
  timeline s-contacts-7, Mentions, the mailbox s-mailbox-21, the agenda).
- Carrel's task 100: the filter of the notebook (s-notes-1), of the users
  (s-admin-2) and of the phone's notes (phone-8) is a `Filter` — a GET form
  on the server's side, its search icon drawn — not a bare search input.
- The agenda's Day / Week / Month / Range is a `SegFold` of links (s-agenda-1,
  states-3, -6, -13, -15): a switch of views is a move to another address,
  and buttons did nothing without a script; on a phone's bar (640 px and
  less, `SegFold.Phone`) it folds to the current view with a menu of the
  rest, so the bar fits a phone with no help from the script (`TestBarSafetyQuiet` was red at 390 px at
  `bc575a8`). 393 frames.
- The new event says «No status» and «Does not repeat», the words of its
  lists (s-agenda-15, Carrel's task 106).
- Carrel's task 94 and the owner's decisions of 02.10: the external trigger
  of a backup job has its Gitea tab — a Gitea Actions workflow, the password
  in the repository's secrets under the name the file gives, the archive of a
  job that downloads kept with `upload-artifact` (s-backup-10) — and its curl
  tab, the bare command to copy (s-backup-11); «Regenerate» is «New password»,
  its hint saying the old one stops working (s-backup-2, -3, -5). 392 frames.
- Carrel's task 93 and the owner's decisions of 02.10: an event's save that
  lost properties has its frame, as a contact's (states-17); an attachment's
  name is a link in every card — in the record it downloads, in Files it
  opens there (`DocAttach.Href`); Remove stands on the edit only
  (s-agenda-3). 390 frames.
- Carrel's tasks 78–82 and the owner's decisions of 02.10: no backup jobs —
  no backups line in the rail and no count by Backup (s-backup-6, 78); the
  invitation by link has its own frame (s-auth-14, 79), and both invitations
  show the address and let it change, another one confirmed by a letter
  (s-auth-12, -14); the states of contacts have whole heads, and a refused
  source shows in the rail too (states-7…-11, 80); Anna's data is the same in
  every edit (s-contacts-4, states-12, 81); fields in a form's columns stand
  by their tops (a birthday with its hint no longer floats up). 387 frames.
- The attendees of an event (Carrel's task 77, the owner's decisions of
  01.10): in the card a list in columns — a known person by the colour of
  their book and a link, an unknown one by the address (s-agenda-1, -2,
  phone-3); in the edit a tag field in columns with suggestions from the
  ticked books and «Add ‹address›» (s-agenda-3, -12, -13); an invited person
  sees the same tags without crosses (s-agenda-14).
- The calendar's new frames (Carrel's task 77, as tasks 59 and 66): a new
  event — the edit's form, empty, one reminder by default, no «Kept as…» and
  no Delete (s-agenda-15); the agenda while the poll runs (states-13), with
  nothing ticked (states-14), and one calendar's empty week — one line, not
  seven empty days (states-15). A day without events is not printed: the
  agenda's empty Wednesday goes (s-agenda-1, the owner's decision of 02.10).
  386 frames.
- The calendar's frames by the owner's decisions of 01.10 (Carrel's task 77):
  Previous · Today · Next lead the agenda's bar (links; on a narrow screen
  Events · Tasks due · Notes go under «⋯», and the bar's second rank from
  1100), New event has its arrow, the agenda names both its addresses, rails
  keep the server's order, the organiser is the attendees' first row, an event
  is edited in its own zone (s-agenda-1..3, -8, phone-3); the states of loading
  and of a source that did not answer have the whole rail, head and bar
  (states-3, -6). A new icon: `next`.
- The timeline is where a person takes part, not where they are mentioned:
  the glossary and the captions say so (Carrel's task 70).

### Fixed

- A date range (`Range`, `.tc-range`) is never wider than its place: where
  it does not fit, Show wraps under the fields (Р-65; `components-18` cut
  its button by 1 px at 390). `Range.Wrap` is not needed for that any more.
  Its sample and `s-mailbox-21` end on a date, as `s-contacts-7` does — not
  on a text field «today» as wide as a date.
- A `Dialog` with an `Action` keeps its form when a `PostForm` is one of
  its parts (a footer button's hidden form, as Sign out): that form is
  printed after the dialog's, not in it. Before, the dialog was printed
  plain and its own fields went nowhere.
- A `Busy` button comes back — its words, on, not busy — when the live
  layer puts the answer in a zone and the form stays; sent again the plain
  way it keeps its own words; a `Folds` or `Bold` label says the busy words
  too.
- A `Field.Changed` brought hidden sends nothing from the start, not only
  after its first show; a control the page brought off (`Check.Reason`)
  stays off when its field shows — by `Changed` or by `PickOption.Shows`.
- A `Seg` of links or of options that send a form (`SegOption.Form`) is a
  `role="group"`: they are no radios, and a `radiogroup` of them gave a
  reader no radio to pick. Plain options and a `Seg` with a `Name` stay a
  `radiogroup`.
- An answer's page forms (`Fragment.Forms`, `ZonePart.Forms`) come in a
  `<template data-tc-forms>` of their own: an answer of forms only now
  empties its zone, as an empty one does, and a zone's content that is
  itself a `PostForm` stays in the zone. An answer printed by hand puts
  its page forms in that template too.

## typecase/v0.11.0 — 2026-10-01

### Breaking

- **`Head.Tools` is `[]HeadTool`** — a `Button` or a `Drop` — instead of
  `[]Button`: a panel's «⋯» stands among its tools with its menu.

### Added

- **`shape` knows the meaning of HTML everywhere.** An element with classes
  is compared by its classes, not its tag, so every link, form, checkbox and
  select the widgets print for their meaning stands for the frame's element;
  a select's insides are the browser's options. `TestMeaningIsNoDifference`
  checks every demo screen with every meaning field set. A form's tick keeps
  the frame's `is-off` and its `:checked` state overrides it.
- **`frame2go -self` output is gofmt-clean** (no blank line at the end).
- **A badge in a sheet's head** (`Badge` is a `SheetHeadPart`): «local» on the
  publish sheet.
- **Columns for every list** (Carrel's task 45). `Column` gains the kinds the
  other lists need — `ColumnDue`, `ColumnComplete`, `ColumnPriority`,
  `ColumnStatus`, `ColumnExcerpt`, `ColumnDate`, `ColumnDetail` (why a record
  is in a timeline) and `ColumnKind` — each with its width and the step it goes
  at on a narrow screen; `ColumnRow` takes a task's `Tick` and `Overdue` (only
  the due is in the alert colour).
- **`Mention`**: a mention in running text whose card is not in the books a
  person looks in — the name, grey and dotted, with a hint (Mentions,
  deferred).

- **`SegFold`** — a crowded switch of links (the owner's decision of 30.09,
  variant B): on a wide bar a segment, on a narrow one — where the bar hides
  what is second (`is-2nd`) — a pick of the current option with a menu of the
  rest, opened by `details` without a script. `SegOption.Href` and `Count`: an
  option may be a link with its number (Carrel's task 50).

- **`shape` reads a picture in an avatar or a thumbnail as data** — the frame
  draws the place, the screen fills it (`Src`); an avatar's `is-photo` is data
  too (Carrel's task 48).
- **Closed menus** (Carrel's task 49): a button with its menu (`Drop`,
  `TopDrop`) is, for `shape`, its button — the frames draw the menu's contents
  beside the screen. `Drop` stands in a head's actions (`Acts`) and among its
  tools: `Head.Tools` is `[]HeadTool` (breaking), `Button` or `Drop`.
- **`Range.Action`**: a bar's range is a GET form — From and To go as fields,
  Show submits it, no script (Carrel's task 51).
- **`FieldsItem.Href`**: a card's value goes somewhere — `mailto:`, `tel:`, the
  linked record; it looks like the value, underlined under the pointer, and
  `shape` reads the link as the value (Carrel's task 52).

- `TestFramesLadder` measures a child from the nearest box above when its
  parent is `display:contents` (the phone bar's navs).

### Mockup

- Sign-in with a wrong login or password (s-auth-13): the same page, the
  password field empty and in focus, one reason for both cases.
- Carrel's tasks 40–43: one page for any dead secret link, the same words and
  one door to sign in; the note to the mockup's reader leaves the recovery
  card for the caption; the publish sheet says the instance is local; a label
  another client wrote is chosen on its own row. The erase page names what goes (the
  encrypted: connections, links, devices, triggers, the deposited key) and what
  stays (login, role, address, first day of the week, the clock) — task 44.
- **Frames of chapters 8–12 lose the `m-` prefix** (the owner's decision):
  `m-states-N` → `states-N`, and so `dark-`, `phone-`, `print-`, `desktop-`,
  like `components-`, `account-` and `rest-`. `m-` was the mockup's own mark
  from before the library; `tc-` stays the prefix of classes and variables
  only, and frame names are not classes. Go names of the screens are the same
  (`Desktop3`). The prose's last `.m-*` class names are the `tc-*` ones from
  `rename.tsv`, and a dead link to `#m-io` points at 7.9.
- The comparison takes `TC_REFACTOR_TAB` (seconds) for a slow machine; the
  whole run on 30.09 was clean — the one difference, s-dup-3, is the frame
  label's «Меню» tag.
- Columns in every list where the button stands (the owner's decisions of
  29–30.09): the menu of tasks (s-tasks-12: Priority, Status and Changed can be
  added), notes (s-notes-23: Changed), files (s-files-20), users (s-admin-21:
  Address), the audit log (s-admin-23, and its bar has the button now) and a
  person's timeline (s-contacts-25: why a record is there stays with its
  title); lists with chosen columns — tasks in one list sorted by Recently
  changed (s-tasks-13), notes (s-notes-24), users with Address (s-admin-22),
  the timeline with the date last (s-contacts-26). Source is in the menu of
  the combined view only. Sorting follows the columns: a sort shows its column.
  The rules under «Выбор столбцов» say where the choice is kept (the server's
  cookie per section), how columns go on a narrow screen and how sorting
  follows. 335 frames.
- **Mentions are approved and deferred** (the owner's decisions of 30.09):
  built after Carrel's first working version; section 7.23 draws them to the
  decisions, and its frames carry the mark «отложена» — neither routes nor
  catalog, like a candidate's. The card folds rare fields under More fields
  (mentions-1, opened in mentions-6; in the form, mentions-7): Related with
  its kind from RFC 6350 (mentions-8), Nickname, Anniversary, Gender,
  Language, Time zone, Geo, Role; messengers stay as the server sent them. A
  link to a card not in the books looked in keeps its name or UID, grey and
  not a link (mentions-4, -6, -7); a server that loses `RELATED` is named
  (mentions-5); the screen has Columns like the timeline (mentions-2). The
  card's timeline has no matches by name any more — the server has none
  (s-contacts-6, -16). The group card (mentions-9, a candidate): members from
  `MEMBER`, found in the books looked in. 339 frames.
- The phone's bottom bar: the sections and «Back to list» are items of one
  bar, all of one width, the way back last behind a hairline — as in the
  rail; before, the two halves split the bar and the way back stood alone on
  the right (the owner's decision of 30.09: the button goes down with the
  rail).
- **`NavItem.Short`**: the label in the phone's bottom bar when the rail's is
  too long for it — «Back to list» in the rail, «Back» in the bar (the
  owner's decision: nothing to save at full width).
- **Mail of one's own mailbox is approved and deferred** (the owner's
  decisions of 30.09) and moves into the mockup as section 7.24, its 32 frames
  (`s-mailbox-N`) marked «отложена»; `carrel-mailbox-mockups.html` is gone. Its
  rules join `data-tc`, its icons the sprite (clip, eye, flag, forward, inbox,
  mailopen, reply, replyall, send, spam); the old davloom mark and rail note
  are the mockup's `tc-emblem` and `tc-railstatus`. Decided: a folder scrolls
  on with «Show 50 more» instead of pages (s-mailbox-1); from a letter the way
  back is «← Inbox» first, with small ↑ ↓ to its neighbours (s-mailbox-6); the
  unread sum, Mail seventh in the phone bar and a dialog of its own to connect
  a mailbox stay as drawn. 371 frames.
- **Deferred frames** — a function approved and put off (mark «отложена») —
  need no reference record, widget or nesting pair until it is taken into
  work (`fragment.Mockup.Built`); lint and the before/after comparison read
  them as all frames. Lint does not read a letter's original in `srcdoc`.
- The filter by kind folds instead of going under «⋯» on a narrow bar: a
  person's timeline (s-contacts-7), Mentions (mentions-2) and the timeline with
  letters (s-mailbox-21); its options are links.
- The contact card's «⋯» and the card's in a timeline are drawn closed with
  their menus (s-contacts-6, -7); a card's phone and address are links
  (s-contacts-6); the bar's range of the timeline and the agenda is a form
  (s-contacts-7, s-agenda-1). A calendar the server made with less than was
  asked for says so over its lists, with «New calendar» (s-agenda-11, Carrel's
  task 53). 372 frames.
- Carrel's tasks 54–57 (the owner's decisions of 01.10): every rail has its
  foot — davloom's mark and the build line (C-13), on the card, the edit, the
  wizards and the conflicts too (48 frames); the rails of a read-only book
  (states-4) and of the full new-contact form (s-contacts-4) carry the whole
  set of books; a finished list has no «Loading more…» (s-contacts-1, -2); the
  quick contact sheet's address is `/new/quick` (s-notes-8).
- Every phone and every address of a card is a link (`tel:`, `mailto:`),
  not only the first (s-contacts-6, Carrel's task 58).
- The full new-contact form has its frame (s-contacts-27, `/new`): the edit's
  form, empty, «Back to list» and «New contact» — the card does not exist yet
  (Carrel's task 59). 373 frames.

## typecase/v0.10.0 — 2026-09-29

### Breaking

- **`AuthStack.Links` are `[]Link`** — text and address — instead of
  `[]string`: «Forgotten your password?» goes to `/forgot`.
- **A menu's heading is a row, not a control**: `RubricItem` prints a `div`
  with `role="presentation"` (the frames draw it so too).

### Added

- **Meaning for everything that is pressed** (task 39 of Carrel's server).
  `Form.ID` — a button outside the form (`Button.Form`) submits it, and Enter
  in a field does the same; `Send` (`Form`, `Name`, `Value`) on every tick
  box — a submit button that acts at once (a source of the rail, a card's
  choice) or a checkbox of the form it stands in (rows before a bulk
  action): `TickRow`, `SizeRow`, `SlimRow`, `TickBox`, `CardPick`, the
  table's head and picks; `Tick.Send` — a task done without a script;
  `Search.Action` and `Filter.Action` — forms of GET, `Query.Name`;
  `UserMenu.Href` — the account's page without a script; `DocRelated.Href`,
  `DocFile.Href`, `Span.Href` — words in the accent that go somewhere.
  Frame components-43 draws each once.
- **The check that finds what still has no meaning**:
  `TestEveryControlHasMeaning` fills every meaning field of every reference
  record and reports each button, link, link-like part and field that still
  sends or leads nowhere, with its frame.

## typecase/v0.9.0 — 2026-09-29

### Added

- **The page around a screen** (task 38 of Carrel's server). Package `page`:
  `Document{Title, Lang, Assets, Root, Forms}` prints the head a page of the
  library needs — the theme's script and the stylesheet by address, nothing
  inline, so `script-src 'self'; style-src 'self'` holds — the hidden
  sprite, one of the library's roots (`widget.Root`: a shell, a page, an
  overlay; closed, so there is no place for markup of one's own) and the
  PostForms. `page.Static()` serves the library's files from the binary;
  `page.At(base)` names their addresses.
- **Columns a person chooses** (the rest of task 29). `ColumnRow` is a row of
  the chosen columns (`Column`, `ColumnCell`): the set and the order are the
  person's, each column's width and the step it goes at are its kind's, and
  the columns line up because every row of a list has the same cells. The
  Columns menu is a form (`Menu.Action`, `Hidden`; `MenuItem.Name`, `Value`,
  `Fixed`, `Drag`): its ticked items are sent in their order — ticks and
  Apply work without a script, and with tc.js an item is dragged by its
  handle or moved by Alt+↑ / Alt+↓. Frames components-42, s-contacts-23, -24.

## typecase/v0.8.0 — 2026-09-28

### Added

- **Widgets carry the meaning of HTML** (task 29 of Carrel's server, step 1).
  A button with `Href` is a link, with `Form` it submits a `PostForm` that
  stands anywhere on the page (a form of hidden fields: CSRF), with `To`,
  `Name` and `Value` it tells its form where and which action; `Form` with
  `Action` is a form element with its hidden `Fields`; `Check`, `Seg` and
  `Pick` with `Name` are a checkbox, radio buttons and a `select` that send
  their value and switch without scripts; `Input` takes `ReadOnly` and
  `Required`; rows, menu items and tree rows take `Href`. Each looks as it did:
  the mockup draws them once (components-37…39) and the CSS matches.
- **`shape` checks a client's screen against its frame.** `shape.Parse` reads
  the mockup, `Check(frame, html)` compares the make-up — elements, order,
  classes and states — and not words or data: a run of rows may be longer or
  shorter, a link stands for a button, a select for a pick.
- **Pictures in their places.** `Avatar`, `Thumb` (with `Alt`), the field's
  `Photo`, `PhotoCrop` and `CardQR` take `Src`: the picture fills the place
  the frames draw in a colour, cropped to it (components-40).
- **The rest that screens need without their own markup.** `List.Compact`
  (rows of 30 px; a phone keeps its own), inputs of 16 px under a finger
  (pointer: coarse), a sheet that opens by its address with no script
  (`.tc-scrim:target`, `Close.Href`), `Check.Disables` (a check that switches
  a field off), `Button.Adds` with `Template` (a row added in place; without a
  script the button submits), `Seg.Theme` (the theme kept in this browser),
  `Scrim.Drop` with `PostForm.File` (a file dropped on the page is sent). A
  closed scrim is hidden by the CSS too, not only by the script.
- **`frame2go -self`** prints a frame's draft as a Go file of its own that
  builds in the client's module with the widget package alone.

## typecase/v0.7.3 — 2026-09-28

### Added

- **A menu item that cannot be chosen now.** `MenuItem.Reason` switches the
  item off and says why in its hint — someone a list names but who cannot be
  picked, the role of the last administrator.
- **A table in a document's text.** `DocTable` prints a Markdown table: a
  head, rows, each column's alignment (`DocAlign`), bold, links and code in
  cells. Words wrap between words, never inside one; links and code do not
  wrap; what does not fit scrolls inside the table. In a container of 640 px
  and less a Wrap / One line switch stands over it (`DocTableSwitch`,
  `SegCompactSmall`), and tc.js keeps the rows to one line when asked.

### Mockup

- Frames for the tasks of Carrel's server 26–28 and 30–34, with the owner's
  decisions: a table in a note wraps between words, and on a narrow screen
  has a Wrap / One line switch; the time zone row stands for every timed
  event; «Tags» everywhere; a backup job with its external trigger off has
  no card for the line; the reset sheet asks for the login. New frames:
  the storage's own row menu, a storage folder's rename, delete and new
  folder, a new collection with no server answering, the password reset's
  answer, Whose and its refusals, the created user's password, the cards
  without mail, field refusals, the invitation page, Related to outside the
  loaded records, the escrow page with the scheme off, New folder's refusal
  in the upload sheet, Browse of the attachment folder. The contact's edit
  has its addresses, the event's edit a second reminder.

## typecase/v0.7.2 — 2026-09-28

### Fixed

- **Date and time fields keep their column on Android.** Mobile Chromium
  drew `input[type=date|time]` in its own look, wider than its field and with
  its arrow at the border; such inputs have no system look now and fill their
  field. Desktop browsers draw them as before.
- **A narrow form fills its width.** In a container of 640 px and less a
  form's field groups and its rows of fields and buttons are as wide as the
  form, instead of as wide as their content and centred — Attach and the
  form's buttons no longer drift to the middle.

### Mockup

- Frames for the tasks of Carrel's server note of 28.09 (8–25), with the
  owner's decisions: deleting is one button with no line under it; Attach
  stands wherever a record is shown; a task has one «All day»; a user's role
  is one item that switches it; the files head's «⋯» is Print only;
  publishing a collection lives in Connections only. New frames: the sheet
  that resets a password, the Related to list, the account's email change
  with no letter gone, withdrawing a deposited key and the escrow page
  without a copy or with the scheme off, the copy sheets, a move to another
  storage, New folder with its name and refusal, From Files and its
  refusals, a read-only folder's Upload, a connection without a password and
  its certificate fingerprint, the rewrite sheet for one record, for nothing
  and its outcome, a public page of one record.

## typecase/v0.7.1 — 2026-09-27

### Added

- **`Input` takes a password and can be switched off.** `InputPassword` is
  the password kind; `Input.Disabled` prints a field that is there but does
  not take input — the connection form's password when the server asks none.
- **A table row keeps its buttons while there is room.** `TableCell.More` is
  the row's «⋯»: it shows only when the row's `is-2nd` buttons are hidden,
  which a table does in a container of 640 px and less.
- **`TimeRow.Short`**: the lead is a time of day only, and its column is as
  narrow as the time — an event among the search results, whose title then
  stands level with a contact's name.

### Mockup

- Frames for the tasks Carrel's server found (plan 4): the note frames tick
  «Дневник» as the list does; the timeline's «Where to look» shows every
  calendar with its tick, feeds too; the search's rail is every source of
  the set; published links are newest first and keep Copy, Send and Revoke
  on the row. New frames: the conflict of an event, a task and a note; the
  sheet after a card's address is replaced, which offers to rewrite it in
  the records where the person attends, and its outcome on the card; the
  connection form of a server that asks no password.
- Row actions by how often they are needed: a backup job keeps Run and Edit
  on its row and is deleted on its page, with no «⋯»; an invite keeps Copy
  link, and Resend where it went by mail, with Extend and Revoke under «⋯»;
  a user's actions are all rare and stay under «⋯». The job row's menu frame
  is gone.

## typecase/v0.7.0 — 2026-09-27

### Breaking — 0.7.0: the library prints the mockup

The widget API is replaced, not extended. Each widget prints the fragment of
its mockup frame — the same tags, classes and text — and the library's CSS is
the mockup's `<style data-tc>` block cut out verbatim into `css/typecase.css`.
What a container accepts is derived from the mockup's frames and fixed by Go
types: a nesting the mockup does not draw does not compile. The `layout`
package, the `[]Node` "any widget" slots and the screen-side fallback
(`Cols`) are gone. Before 1.0 a breaking change comes in a minor release
without a deprecation period; there are no aliases or transition layers, and
this entry says what changed rather than how to migrate. How a screen is
built now — a frame of the mockup, drafted by `cmd/frame2go`, its data then
the application's — is [docs/guide/screen.md](docs/guide/screen.md). The spec
(`docs/spec/typecase.md`) states the model: L-1, L-2.2, L-3.1, L-3.2 (the
eighth item is now «Эталон», the frame and fragment a widget prints), L-3.6,
L-5, L-6.1, L-7, L-10, L-11.1.

- **One stylesheet.** `css/typecase.css` is the mockup's `<style data-tc>`
  block, cut verbatim by `go generate ./css` (`internal/mockcss/cut`); the
  gate `css-from-mockup` compares them byte for byte. `css.Order` is
  `fonts.css`, `typecase.css`; the eleven hand-written `css/*.css` files and
  the generated layout CSS are gone. `token` requires what the mockup
  defines — the palette, three font families, the page gutters and the layout
  tokens `--tc-rail`, `--tc-side`, `--tc-top` — and rejects a token that is
  read but never defined; `--tc-font-size`, `--tc-row-height`,
  `--tc-radius*` and `--tc-stroke` no longer exist. `tc.js` sets
  `--tc-colour` and `--tc-width` from `data-tc-colour` and `data-tc-width`
  and the mockup's CSS applies them; `data-tc-pct` is gone. The mockup's
  rename, element and inline scripts are removed.
- **A widget is checked against its frame.** `internal/fragment` parses the
  mockup (`golang.org/x/net/html`, internal only), finds a block by
  `Ref{Frame, Class, N}`, normalizes both trees — tags, classes, text, field
  values, icons, `data-tc-colour`/`data-tc-width`; not `aria-*`, `role`,
  `id`, `title` — and names the first path where they part.
  `go run ./internal/fragment/cmd/shapes tc-tickrow` prints a block's forms.
  `internal/refs` holds the reference table and the list of widget types;
  `TestWidgetFragments` compares every record.
- **What a container takes comes from the mockup.** A container declares its
  children by field types: a widget field is a fixed child, a slot is an
  interface with an unexported method that exactly the widgets the mockup
  puts there implement. `TestContainerPairs` derives every parent → child
  block pair from the frames (362 today) and checks them against the types
  both ways; an interface field that takes any widget is refused. How the
  library reads a frame — the block of an element, mixins such as
  `tc-push` or `is-2nd` that the container prints, forms without slot
  contents — is written in `docs/visual/mockup.md`.
- **Primitives.** `Button`, `Close`, `Mark`, `Tag`, `Badge`, `Tick`, `Kbd`,
  `Num`, `Avatar`, `Dim`, `Mono`, `Plain` print their frames' fragments,
  checked form by form (124 records). Kinds, states and variants are enums;
  running text inside `Dim` is `Text`, `Strong`, `Hint`, `Dot` and `Mono`.
  `Icon` names a symbol of the sprite without enumerating the set, and
  `icons/sprite.svg` is now the mockup's sprite, cut like the CSS. The old
  widgets live in `internal/legacy/widget` until step 6.
- **Fields and forms.** `Input`, `Pick`, `Filter`, `TagInput` and `Area` are
  the five forms of `tc-in` (one type each, so a pick with a date's
  placeholder cannot be built); `Query`, `Seg`, `Check`, `Range`, `Field` with
  its parts (hint, swatches, zoom, row, group, photo), `FieldSet`, `Form`,
  `FormFoot`, `Rubric`, and the layout glue `Line`, `Stack`, `Grow`, `Push`.
  `Place` (pushed, grown) and `Secondary` (is-2nd) are fields of the widgets
  the frames place so. Forms whose children come from later families wait in
  `refs.Waiting` with their reason; `TestShapesCovered` holds every other form
  to a record. In the mockup: the zoom is `tc-field__zoom`, the weights
  column `tc-rubric__col`, and «Find a zone» gets the `is-ph` every other
  placeholder has (awaiting the owner in `allowed-diffs.md`).
- **Rows and lists.** A list row is typed by its columns: `TimeRow`,
  `FaceRow`, `ExcerptRow`, `DueRow`, `FindRow`, `LinkedRow`, `BriefRow`,
  `SizeRow`, `SlimRow`, `PhoneRow` and the loading `SkelRow` are the nine
  column sets of `tc-row`, so a row of times has no avatar and a row of people
  no due date. `List` holds any of them; `TickRow` is the rail's row with its
  tick box (`Ticked`: on, part, off), also in a group's grid of choices;
  `Tree` with its rows names the chosen one by number, so two cannot be
  chosen; `Skel` is the placeholder. In the mockup the tree's icons lose the
  four different `svg` wrappers that drew the same icon and keep the plain one
  the tree's CSS sizes.
- **The shell's pieces.** `Top` with `Brand` (a `Logo` or a section's title),
  `Search`, `UserMenu`; `Nav` names its current section by number;
  `RailSec`, `RailFoot` with `RailStatus` (the dot's colour follows the
  state) and the build line; `Emblem` is a service's three-part mark from the
  sprite (`#<symbol>-stroke`, `-fill`, `-accent`); `Crumbs`, `Acts`, `Split`,
  `Zone`. Renamed in the mockup and the CSS: `tc-davmark` is `tc-emblem`
  (parts `__stroke`, `__fill`, `__accent`), `tc-railnote` is
  `tc-railstatus`. The desktop window's title bar and status line are page
  chrome, read past by the fragment reader; its classes that sat in the
  product's namespace (`tc-win--*`, `tc-win__bar`, `tc-winbtn`,
  `tc-win__toast`, `tc-win__title`, `tc-tray__row`) are `mk-*` in the page's
  style now. `Shell` holds the top bar and a `ShellBody` (the rail's pieces,
  `Main`, an optional `Side`); `Head` has one title with the lines, the path
  and the actions around it; `Bar` keeps one row, its secondary items go
  under «⋯» first. `Main` takes the page's pieces, among them the list
  group's heading (`ListGroup`) and a public page's banner, address strip and
  foot. The badge check looks at the badge element itself, so a badge beside
  a link no longer fails it.
- **Overlays.** `Menu` with its items (`MenuItem`: icon, mark, tick box,
  label or found words, a monospaced detail, a trail), separators, headings,
  value rows, a column with a switch and a foot; `Drop` is a button or a pick
  with its menu, and is its bar's «⋯» when it says so; the top bar takes a
  menu under the search and buttons with menus (`TopDrop`). `Sheet` has one
  bold title after a toned icon, what stands at the end of its head, a body
  of named parts (fields, keys, a picker, checks, a block of lines, the
  sender) and a foot; `Scrim` holds it; `Overlay` holds the shell's body
  under the scrim with its own `OverlayBody`, because the side panel is
  closed while a sheet is open. `Dialog` is the centred canvas with its own
  title, logo, tagline and sign-in column; `Auth` is the sign-in page. In the
  mockup the accent icons of sheet heads share one stroke (1.4, round caps)
  instead of three, and the folder picker of «Attach a file» gets the base
  class its sibling has.
- **The document.** `Prose` keeps its paragraphs as written (the text is
  pre-wrapped), with bold, monospaced pieces and links (`Link`) inside; `Doc`
  is a note as read, edited, published, rendered or on a phone, made of named
  parts (title, date line, paragraphs, tasks with real ticks, attachments,
  related, the editor's fields and note, headings, bullets); `Docbar` is the
  edit bar with its end; `DocLayout` puts them in the page, and `Hero` heads
  a service's own page. In the mockup the davloom tile of that page is the
  sprite's symbol `davloom-tile` instead of an inline drawing — the same
  picture, referenced like the logo.
- **Tables, cards and the rest.** `Table` takes columns (width, rank — the
  third and fourth columns go first on a narrow screen — and whether it is
  secondary) and rows of cells; a cell takes its column's rank and tone, and
  a table without a head carries them itself. `Card` is made of named parts:
  lines of text by kind, hints, rows of buttons and fields, lists, a code to
  scan, a choice, a grid, rows of requests, a letter's head. `Sec` and
  `Fields` fill panels; `Msg`, `Diag`, `Danger`, `Meter`, `Runbar`, `Queue`,
  `Stats`, `Steps`, `Diff`, `Empty`, `Thumb`, `Tile`, `Grid`, `Pad` and the
  document's side column `Docmeta` complete the set. Every form of every
  block the frames draw has its record; one waits for the owner (an input
  inside a check's button). The catalog rule for table columns follows the
  frames now: a row may not be wider than the head (a spanning cell counts as
  its columns), and a table without a head or a row that stops short is
  allowed. In the mockup the icons of message rows and of grid thumbnails
  lose a wrapper that repeated their block's rule, and the rule that only
  that wrapper used is gone.
- **tc.js on the new markup.** The script reads what the widgets print:
  a drop's button is the first child of `.tc-drop` (`.tc-drop__anchor` in the
  top bar) and its menu stands beside it; a sheet or a panel closes by its
  `.tc-close`, by a click on the scrim or by Esc, the topmost layer first; a
  control with `data-tc-open="id"` opens the sheet or panel with that id.
  Closed is the `hidden` attribute (backed by an inline `display:none` where a
  rule of the mockup would show the element), since the mockup's CSS has no
  closed state; `Drop` and `TopDrop` have `Open`. The shell's state is read
  from `--tc-shell-mode` and written on the shell as `data-tc-mode`; the
  script keeps no threshold of its own. A bar the CSS did not fit is marked
  `data-tc-overflow` and scrolls sideways. The old hooks (`data-tc-menu`,
  `data-tc-bar-*`, `data-tc-top-*`, `data-tc-seg*`, `data-tc-tree*`,
  `data-tc-pick*`, `data-tc-crop`, `data-tc-rail*` and the rest) are gone with
  the markup they served.
- **`layout` is gone, and the gates are what the library checks now.** The
  `layout` package (tables, generation, inventory, parity) is removed, with
  the old widgets and the old demo screens that stood on it (`cmd/screens`;
  plan 3 rebuilds the screens from the frames). Gates removed: `layout-slots`,
  `mockup-parity`, `layout-inventory`, `layout-orphans`, `geometry`,
  `completeness`, `screens-coverage`, `screens-compose`. `css-from-mockup`
  also refuses a rule outside `css/` that styles a widget; `node-names` reads
  the output of every reference record and wants a library class on every
  element the mockup's dictionary does not leave bare. Spec L-8 lists the 17
  gates as they are. `widget-fragments` also reads the widget types from the source
  and wants each in `refs.Widgets` and with at least one record.
- **The owner's answers of step 7.** A bar gives what did not fit to «⋯» on
  the narrowest steps too, instead of scrolling sideways: the third rank
  `is-3rd` goes under «⋯» at 460 px and below, after `is-2nd`, and its twin
  in the drawn menu of «⋯» shows there only. `Button`, `Check`, `BarSep`,
  `MenuItem` and `MenuSep` take `Tertiary`; a folding button label in a bar
  leaves at the same step, and the fields of a range in a bar shrink. The
  safety of `tc.js` trips on no reference bar at any step. In the rail of
  icons the current section, pressed again, opens the sheet with its
  sources: `Nav.Opens` names it, and elsewhere the item stays a link. A check
  no longer takes an input (`Check.Parts` without `Input`): the input of a
  sentence stands beside the check in a `Line`.

- **Demo screens are the mockup's frames.** `go run ./cmd/screens` serves
  every route frame (decision Р-16) at `/f/<data-frame>`, built from widgets
  with the frame's data, top to bottom; 127 of 154 build, and the other 27 —
  pages without a shell and print — are listed in `docs/visual/gaps.md`. The
  static comparison (`TestFramesStatic`, `internal/fragment`'s
  `CompareFrame`) holds each screen's whole tree to its frame; the new gate
  `frame-coverage` wants every route frame either built or listed as a gap.
  The screens are drafted by `tools/framegen`. What the whole frames showed
  the library lacked: `Pick` in a bar, `Field` in a dialog, `Check` in a
  sheet, a tag input as a drop's trigger, `Place` on `Diag` and `Table`.

- **Each screen is laid out against its frame on every step of the ladder.**
  `TestFramesLadder` puts a frame and its screen on one page and compares
  the box of every node and text run (half a pixel) and the computed style
  that draws it, on all ten steps; it also holds the screen to the
  invariants of a screen (nothing scrolls sideways, siblings do not overlap,
  sheets and menus are on top, text stays in its box, a 44 px zone in the
  drawer state). Ten breakages of the mockup itself on the narrow steps are
  listed in the test with their reasons. Long labels (the safety of `tc.js`
  catches the bars) and a tablet's viewport are run too. What the ladder
  found in the library: `Check`, `FieldHint`, `Faced` and `RubricAside` no
  longer put a space between their parts — the text carries its spaces, as
  in HTML, and a space before a comma is not forced any more; `TableCell`
  takes `Rows` (one value for several rows); the static comparison compares
  `colspan` and `rowspan`. A route carries the theme its frame is drawn in.

- **Pages without a shell.** `Page` is the root of a page that has no top
  bar and rail: sign-in and the password pages, «about», a desktop window's
  inside, a page on paper (variants: plain, print, on the background, on the
  background in a column). `Window` is a desktop window's inside with the
  padding of a big or a small window, `WindowNote` its sentence of purpose;
  a page on paper is `PrintTitle`, `PrintSources`, `PrintRubric`, `PrintRow`
  (a time and what is at it, or a name and two numbers: `PrintCell`,
  `PrintNum`) and `PrintFoot`. A line takes a `Pick`. All 154 route frames
  are screens now and `docs/visual/gaps.md` is empty; a page on paper is also
  compared with the print medium emulated. Reading the mockup, a
  notification of the system (`mk-toast`) is chrome like the window's title
  bar, and a route frame's plain `tc-root` is the page's block.

- **The catalog shows the mockup's forms.** A cell is a reference record, one
  per state; after a widget's last cell stand the mockup's component and
  fragment frames built around it (109; the candidate frames are not shown).
  A switch sets the cells' width on the mockup's ladder, links set the theme.
  An open menu, the search's suggestions and a scrim hang out of their box as
  in their frames, and the catalog makes room for them in the cell. The
  catalog's browser tests run again, on the new markup.
- **A segmented switch chooses.** `tc.js` now does what L-4 says of `tc-seg`:
  pressing an option chooses it (`is-on` and `aria-checked` move to it) and
  the arrow keys move the choice along the switch.

- **A screen starts from its frame: `cmd/frame2go`.** `go run
  ./cmd/frame2go <data-frame>` prints a frame of the mockup as a screen
  function — the widgets and the frame's data, top to bottom, each variant
  read from the widgets' own `class()` tables — that already passes the
  static comparison; the author then replaces the frame's data with the
  application's. Without arguments it writes `internal/screens` (`go
  generate ./internal/screens`), and `TestGenerated` wants the demo screens
  to be its output as is. It replaces the Python scripts of
  `tools/framegen`, which are gone. How to build a screen:
  [docs/guide/screen.md](docs/guide/screen.md).

### Mockup

- Breakages on the narrow steps that the ladder comparison found, fixed:
  the scrim is a layer of its own above the page (`z-index:7`), so the
  drawer state's raised buttons no longer show over a sheet («Перемещение и
  копирование» [s-files-4], «Учётная запись» [account-6], a publishing
  sheet [s-publish-5] at 390 px); with the search folded into its icon, its
  suggestions span the top bar instead of the icon's width («Подсказки»
  [s-find-2], [s-find-3]); a primary button drawn as an icon hides its label
  instead of cutting it («Конфликт · планшет» [s-conflict-2], «Заметка ·
  чтение и правка» [m-phone-7] at 640 px). Nothing moves at 820 px and up.
- Step 7 of plan 2, the owner's answers of 2026-09-26. The bars of «Заметка ·
  правка» [s-notes-20], «Календарь · агенда со списком и карточкой»
  [s-agenda-1], «Администрирование · пользователи» [s-admin-2] and «Человек ·
  хронология» [s-contacts-7] fit at 460 and 390 px: the third rank `is-3rd`
  (Р-4) goes under «⋯» there, the agenda's menu gets the three ticks as its
  twins, Columns keeps its icon, the date range shrinks. In «Резервная копия ·
  задание, что и куда» [s-backup-2] the days input leaves the check's button
  and stands beside it in a line — the frame is the same to the pixel; the
  rule that makes a check beside an input 28 px high leaves the 22 px chip
  input alone. The caption of «Рейл из значков · источники открыты» [m-phone-9]
  says what opens the sheet: the current section's icon, pressed again. The
  placeholder of «Find a zone» and the one stroke of the sheet heads' accent
  icons are the owner's, in `allowed-diffs.md`.
- Mailbox candidate (`carrel-mailbox-mockups.html`): each mailbox has its own
  signature and name in From, set in a new «Ящик · правка» frame and shown in
  the reply sheet under the text, after «-- ». The frame of a separate search
  in the mail section is gone: the header search is global everywhere, and
  mail is narrowed by the Mail segment and the folder ticks.
- The mockup's file rows keep the icon on the name's line, as the library's
  `css/06-lists.css` already does: `.tc-row__name` is a row there too, and
  `.tc-row__file` cuts a long name with an ellipsis. Before, the mockup's
  `.tc-root svg{display:block}` put the icon above the name in «Файлы · обзор
  папки» and «Поиск по всем источникам» (recorded in `allowed-diffs.md`). Nothing
  changes for consumers; the gap is 7 px in the mockup and 0.35em in the
  library, and plan 2 takes the mockup's.
- On a phone (`--tc-shell-mode: drawer`) a button next to a field — Browse,
  Copy, Change, Send test — is 38 px like the field, not 28: the drawer rule
  now lists `.tc-form`, `.tc-fset` and `.tc-field` buttons (small `is-sm`
  ones keep their size). Only widths of 640 and less change.
- Mailbox candidate: Cc and Bcc stand next to the To field as ordinary
  buttons, the «button next to a field» component, instead of inside its box.
- A mailbox gets a colour when connected, as a collection does: its mark in
  the rail and the mark of its letters in search and the timeline take it;
  `.tc-mark.is-mail`, the hollow mark, is gone. Cc and Bcc in the letter sheet
  are hidden behind buttons at the end of the To line and shown once pressed
  or filled (new frame «Новое письмо · Cc и Bcc спрятаны», 32 in all).
- The candidate file lives only in Typecase; Carrel no longer has a copy.
- The mockup's home is Typecase: it is edited here and copied whole into
  Carrel (`docs/visual/mockup.md`), the reverse of the earlier order.
- A menu stands under its own button everywhere: right edge to the button's
  right edge, 4 px below. `.tc-drop.is-bar`, which measured the bar's
  overflow menu from the bar and so shifted it by the bar's padding, is gone;
  the bar's «⋯» is an ordinary `.tc-drop`, hidden whole while it is not
  needed. The overflow menu has no «Didn't fit» rubric.
- Nothing scrolls sideways inside a frame. `.tc-bar--wide` (agenda, person
  timeline) sends `is-2nd` under «⋯» from 960 px, `.tc-table--wide` hides its
  `is-2nd` columns from 800; the note edit bar gives up `is-2nd` from 1120
  (it needs about 1115, not 1008) and its button labels from 740;
  `.tc-bar__right` labels do not wrap; `.tc-fset` and its grid may shrink;
  header actions are never wider than the header. A table row shows one
  action in words and the rest as an icon or under «⋯»; new row menus are
  drawn by their screens.
- A disabled button keeps its colour and fades (`.tc-btn:disabled`); why it is
  disabled is a line under the button row (`.tc-card__why`), not beside it.
  Remove in a form and Revoke are `is-danger`.
- Moving a contact or files is a sheet over the list, like moving a calendar.
  New frames: «Заметка · правка» (the note page in edit, with the rail and
  the column; Focus is its state), account states (address confirmed,
  confirmation sent, key deposited). Name and Reply-To live on the one
  account screen; the separate account frame in the mail section is now its
  «not confirmed» state beside it. Section 7.4 says why notes differ: the
  panel edits, and the note page lists its neighbours.
- Birthday is typed, not picked: a text field, the year optional. The
  duplicates weights name both columns, Default and Yours. The public note
  page offers only the download the author allowed, `.md`; the files header
  «⋯» no longer repeats the rail's Published links.
- The width ladder (`TestMockupLadder`) now fails on horizontal scroll of any
  box inside a frame, not only of the frame: a scrolling `.tc-main` had hidden
  a too-wide bar.
- The note panel is a view, as for the other kinds: formatted text with its
  tasks, tags, attachments and links, with Edit, the note page and «⋯» under
  the title. Editing is on the note page («Заметка · правка»); section 7.4
  now names the one real difference, neighbours in the note page's rail.
- Focus has its own icon, `i-focus` (viewfinder corners). The outward
  arrows, `i-expand`, mean one thing across the mockup: open larger — the
  note page from the panel, the card from a «New …» sheet.

- The owner looked through the 69 frames the one-off fold moved and accepted
  them; the screenshot base (`docs/visual/allowed-diffs.md`) moved to
  `de784a4` and its table is empty. `docs/plan/fold-review.md` is gone.

- Plan 1 (mockup tails) is closed and deleted; what lasts from it —
  `tc-plain`, the fold of one-off classes, states decided not to exist, the
  print footer without the mark — is in `docs/visual/mockup.md`.

Next: the breaking 0.7.0 where the library prints the mockup ([docs/plan/2-library-from-mockup.md](docs/plan/2-library-from-mockup.md)).

## typecase/v0.6.2 — 2026-09-25

One thing for consumers: the wide details panel is gone, and every panel —
including a full-screen note's side column — is 400 px (see below). The
removal is breaking and ships as a patch for the same reason as 0.6.1: 0.7.0 is
the release of plan 2, and before 1.0 a breaking change needs no deprecation
period (L-10). Everything else is the mockup: menus and dropdowns drawn next to
their screens, and a separate file of candidate frames for the mailbox.

### Breaking

- **The wide details panel is gone: `widget.SideSizeWide`, the `SideSize`
  type, the `Side.Size` field and the `.tc-side--wide` rule (440 px).** The
  mockup has one panel width, 400 px, on every screen. Drop `Size:` from your
  `widget.Side` literals; a panel that was wide is now 400 px like the others.

### Changed

- **The side column of a full-screen note is the details panel: 400 px, not
  296.** `.tc-doc__layout` is `minmax(0,1fr) 400px` and `.tc-docmeta` is
  400 px wide, as wide as `.tc-side` on every list; in a work area of 880 px
  or less (`@container tcbox`) the column narrows to 296 again, and at 640 it
  drops under the text as before. Nothing to do unless you sized something to
  the old 296 — then measure against the new column. `.tc-cols--sidebar` keeps
  296: it is a page column, not the details panel, and no longer follows the
  note's column.

### Mockup

- A full-width field owns its line; only its control stops at 760 px — the
  due date no longer climbs next to the description. The task form takes the
  event's rubrics, head and foot; the published-links list creates nothing and
  lists no feeds (a link is made where the thing lies); the bar's overflow
  menu is open only while its "⋯" is shown; every details panel is 400 px; the
  publish sheets fit their frame and keep the switch and the term at their own
  width.
- Edit forms keep one Save/Cancel pair, in the head; the foot holds only the
  deletion. Four "⋯" with nothing behind them on their screen are gone
  (timeline head, backup job rows, the task bar, the link rows under the
  letter), and the task card in the list panel gets the "⋯" the other cards
  have. Every remaining "⋯" and button arrow on a screen is drawn open as its
  own frame right after that screen — 19 new frames, 226 in all.
- Attach carries its arrow everywhere a record takes a file (from this
  computer or from Files). Every dropdown field on a screen has its options
  drawn as a frame next to that screen — statuses, priority, reminders, time
  zone, repeat, phone and email labels, target collections, link term,
  appearance, admin roles, mail encryption, the audit filter, the desktop
  settings — 33 new frames, 259 in all. The appearance time zone is shown
  read-only: it is the instance's, set by the administrator.
- **Candidate frames for the mailbox: `docs/visual/carrel-mailbox-mockups.html`,
  a separate file.** The deferred feature of Carrel's `carrel-mailbox.md`
  (a person's own mailbox over IMAP and SMTP), in 31 frames: the mail section
  and folder list, the letter, the original in a sandbox with external
  pictures behind a button, mail in the general search and in a person's
  timeline, the letter sheet sending from one's own mailbox, connecting a
  mailbox, the connections and instance settings screens, confirmations.
  **Nothing here is for the library yet.** The file copies the main
  mockup's CSS and sprite as of Carrel `e9c5180` and keeps its own new rules
  in a third block, `<style data-tc-candidate>` (`.tc-row--mailcols`,
  `.tc-mark.is-mail`, `.tc-quote`, `.tc-original`, a bulleted list); the
  main mockup, its lint and its checks do not read it. If the feature is taken,
  the section moves into `carrel-ui-mockups.html` and the block into
  `data-tc`. Its home is Carrel, like the main mockup's; the copy here is
  byte-identical.

## typecase/v0.6.1 — 2026-09-24

Two things for consumers: `layout.ShellNoRailHasSide` is gone, and the source
column of the grid rows changes width and leaves on a phone (see below). The
removal is breaking and still ships as a patch: 0.7.0 is the release of plan 2
(Р-5), and before 1.0 a breaking change needs no deprecation period (L-10).
Everything else is the mockup: its tails are closed
([docs/plan/1-mockup-tails.md](docs/plan/1-mockup-tails.md)), its frames now
agree with the Carrel server or say what the server has to change
(Carrel `docs/plan/from-typecase.md`), and its one-off classes are folded,
waiting for the owner's eye.

### Breaking

- **`layout.ShellNoRailHasSide` is removed, and with it the rule it generated:**
  `.tc-shell__body:not(:has(> .tc-shell__rail)):has(> .tc-shell__side)
  {display:grid;grid-template-columns:minmax(0,1fr) auto}`. A shell with no
  rail but with a details panel is not a state the product has — asked and
  answered — so the bundle was shipping geometry that could never apply. If you
  reference the variable, drop the reference; if you relied on the rule, you
  were relying on a shell that does not exist.

### Changed

- **The source column of every grid row is sized to real stamps and leaves on
  a phone.** In all six grid rows (`tc-row--avatarcols`, `--timecols`,
  `--tickcols`, `--tickmarkcols`, `--excerpt`, `--find`) the source track is
  140 px on the wide steps and 130 px at the 620 step, up from 92–128 px, and
  the title's minimum gives up the same amount, so no step needs more room
  than before. At the 460 step the source is hidden on all six. Two rows
  change what they keep there: the contact row keeps the phone and drops the
  source (it used to drop the phone and keep a 74 px source), and the file row
  keeps the size and drops the source (the library used to drop the size,
  which the mockup never did). If you print a contact row, its source cell now
  needs the `source` slot like the other five.

### For maintainers (mockup source)

- **Two tools to look at the mockup, and two rules it had only in CSS.**
  `TestMockupFrameShots` (`TC_FRAME_SHOTS=frame@width,…`) writes full-size
  pictures of frames; `TestMockupFrameMoves` (`TC_FRAME_MOVES=<commit>`) lists
  the frames whose boxes moved against a commit, on all ten widths, with the
  first element that moved. Both skip without their variable.
  `docs/visual/mockup.md` now says that a menu with fields ends in buttons
  (`tc-menu__foot`), and that a single choice in a sheet is a tree with a
  marked row, never rail nav items, which fold to icons at 641–1007. The
  row-height paragraph no longer promises a `#check` that is gone.
- **A published link may have no expiry.** "Link lives for" offers 7, 14 or
  30 days and "No expiry" (caption of `s-publish-2`); the file link in the
  list of published links (`s-publish-1`) shows "No expiry".
- **The timeline drops rows found by being written about; the rubric aside
  keeps its phrases whole.** The three rows that were "mentioned in the text /
  title" are gone from the timeline, not relabelled: being written about is
  Mentions. `.tc-rubric__aside` has a rule again — its parts wrap as wholes,
  no phrase breaks and none spills past the card. The notes rail loses the
  notebook "Архив", which the fixture set does not have.
- **One-off value classes are folded (Р-13).** Of the 282 rules whose class
  sits on one node, 109 are gone: 67 replaced by the nearest shared class of
  the same block, 42 dropped to the block's base look, whichever is nearer. The 173 kept are
  elements of one frame, variants with a meaning, and geometry a frame sets
  for itself. The drawing moved in 69 frames; the owner reviewed them by eye and accepted them.
- **A taken connection label is refused, in a frame of its own.** `s-profile-11`
  is the connect dialog after Connect with a label that already exists: the
  field is marked and says to choose another; the other fields stay as typed.
  An empty label is the server's address, so a second label-less connection to
  the same server meets the same refusal. New state `tc-in.is-error` (danger
  border). The davloom rights switch sits at the right edge of its cell.
  Frames are 207.
- **A calendar is attached whole or for dates, from one menu item.** The
  attach menu says "Calendar" instead of "Calendar for a range" (`s-mail-2`),
  and the sheet (`s-mail-10`, "Attach a calendar") asks "Whole calendar" or
  "Dates"; From, To and Show belong to Dates. On a phone the calendar list
  stands above the fields.
- **Device rights are back, and new-device passwords stay on the page.** By
  the owner's decision a davloom device is again "read only" or "read and
  write": chosen when it is added (`s-davloom-2`), shown in the step-2 subtitle
  (`s-davloom-3`) and switched in the device's row (`s-davloom-1`), with no
  proxy-wide ceiling. The new-device screens no longer send the person to
  Connections for a password: Android draws one card per ticked account, its
  code, "Open in DAVx5" and the password with Copy (`s-plan-2`); Thunderbird
  and "Anything else" carry login and password with "Copy password" in the
  account row (`s-plan-4`, `s-plan-5`). The weights menu ends in real buttons,
  Save and "Back to the instance's numbers" (`tc-menu__foot`, `s-dup-3`): a
  menu item acts at once, fields wait for Save. The server side is a note in
  Carrel `docs/plan/from-typecase.md`.
- **The numbered classes are gone (Х-1).** The closed plan's inline script had
  moved every distinct `style=` line into a class with a number: 602 of them
  over 1717 uses, and 192 hung on a tag name (`tc-div--r37`) rather than on a
  block, which Р-6 does not allow and which plan 2 would have had a widget
  print. They are now 596 names — by meaning where there is one
  (`tc-in--body`, `tc-empty--drop`, `tc-card__acts`, `tc-sheet__title`), by
  value where there is not (`tc-dialog--max520`, `tc-mark--h18`), as the file
  already named `tc-table__col--w110` and `tc-line--gap-14`. A variant that
  splits into independent parts is written in parts: `tc-root--on-bg
  tc-root--flat` rather than one long name. New lint rule
  `no-numbered-classes` keeps them out. **Nothing moved on screen** — see the new check below.
- **Three names joined the dictionary**, because without them there was
  nothing to call a nameless `<div>`: `tc-dim` (secondary text, which turned
  up in fifteen blocks and so is a block, with size steps `--t10-5`…`--t12-5`
  and tones `--muted`, `--accent`, `--danger`), `is-focus` (the drawn focus
  ring, which `.tc-in` already had, now also on `.tc-btn`, `.tc-search`,
  `.tc-usermenu`, `.tc-mark`), and `tc-plain` (ordinary case under a heading
  that shouts). The glue variants that looked necessary turned out not to be:
  of the seventeen numbered rules with `display:flex`, none matched `tc-line`,
  so nothing was hiding there.
- **A name may exist without a rule, and it says so out loud (Х-4).** The
  mockup listed nineteen classes in a `.tc-shell,…{}` rule with no
  declarations so they would pass `classes-defined`; such a rule says nothing,
  passes `dead-rules` because its selector matches, and would be cut into the
  library's CSS. Those names now stand in a
  `/* names-without-rule: … */` comment in the sheet they belong to, and the
  empty rules are gone. New lint rule `empty-rules`: no rule anywhere in either
  `<style>` may be empty, and every listed name must stand on a node.
  `classes-defined` reads the list.
- **Frame widths are page classes, and the width switch now reaches every
  frame (Х-5).** Twenty-one frames carried `style="max-width:900px"` and the
  like. Inline styles outrank any selector, so eleven of them ignored the width
  ladder above the contents. They now carry `.frame.w392` / `.frame.cap900` and
  friends, defined next to the switch, which is more specific and wins.
  `frame-no-style` covers the `.frame` element itself from now on, so this
  cannot come back; `.frame-label` stays exempt as page chrome.
  `TestMockupLadderSwitch` checks the eleven in a browser.
- **The dark theme gets the two frames its rules need (Х-2).** Tokens carry it
  almost everywhere, but the avatar and the preview tile are filled with a hard
  `#D9D2BE` and the scrim is ink at .22 over a light page, so each needs a rule
  of its own — and the lock had deleted all three because no dark frame showed
  those components. `m-dark-2` (the drop overlay over the contacts list: avatars
  under a scrim) and `m-dark-3` (files as tiles) are copies of existing frames
  with one attribute on the root, which is what "any frame can be shown dark
  with one attribute" was supposed to mean. `TestMockupDarkRules` checks the
  nodes end up the right colour, which `dead-rules` does not: it only says a
  selector matched something.
- **Five of the other deleted rules are not missing frames at all.** Checked
  against the base commit rather than guessed: `.tc-badge.is-plan` and
  `.tc-table td.is-idle` never stood on a node even then, so they are
  inheritance, not undrawn states — and a human has since confirmed a table row
  has no "idle" state; `.tc-queue.is-min ul` has its frame already
  (`s-files-3` draws the collapsed queue by not drawing the list), so that rule
  belongs to the library, not here; `.tc-table__col--w34` was junk; and the
  no-rail-with-panel shell does not exist, which is what the removal above is.
  `mockup-parity` no longer carries a hole for it.
- **A component is shown by a frame, not by a showcase in the running text
  (Х-3).** The components section had 27 `.demo` blocks that were half inside
  the rules: `dead-rules` counted them as product and kept rules alive, while
  none of the other frame checks applied. A frame is the only unit the library
  is measured against, so intent that sits outside one cannot be checked. 26 of
  the 27 show real product blocks and are now frames `components-9…34`, built
  the way `components-1…8` already were. The 27th is a strip of logo and icon
  sizes: it stays page decoration, and its layout (`tc-specs*` → `specs*`) and
  the 38 px logo size moved to the page sheet, because no frame draws them.
  Frames 167 → 193.
- **`dead-rules` now means "fires in a frame".** The "or in a showcase" escape
  hatch is gone with the showcases; 52 product classes had been leaning on it.
  Nothing moved on screen: the 26 boxes compute the same display, direction,
  wrap, alignment, gap, padding, background and width they got from the `.demo`
  page rule, which is what the three new root variants (`tc-root--wrap`,
  `tc-root--col`, `tc-root--pad18-20`) exist to reproduce.
- **`data-check="off"` is gone** with `data-demo`: it steered the `#check`
  in-page audit, which was a script, and scripts are not allowed in the mockup.
- **`data-demo` is gone.** It listed pre-rename `m-*` names on 34 nodes and
  nothing in either repository read it. The rename history stays in
  `docs/visual/rename.tsv`.
- **The `.tc-queue` showcase is named like the frame** — `tc-queue__list`,
  `tc-queue__name`, `tc-queue__meta` instead of bare `<ul>` and `<span>`.
- **The settings rail has a footer, and it carries the statuses that have to be
  in sight.** That footer was two different things without saying so: on the app
  rails it holds the buttons that add what the rail lists (new notebook, connect
  account, subscribe), while on the settings and admin rails it held whatever
  the one screen happened to have — an SMTP badge, a backup date, a davloom
  icon — and only on 7 frames out of 38. Those rails are short and fixed, so the
  footer is now one component, `tc-railnote`: a status dot (`tc-dot--accent`
  running, `tc-dot--dusty` off, `tc-dot--danger` alarm), a title, and a second
  line of the shape `Last <what> <when>`. The set is fixed per panel and does
  not follow the section you are in — settings show davloom and Backup,
  administration shows mail and the installation checks — because the point is
  not having to know which section to open to see them. Ordinary states are
  shown too, not only breakage: a banner that appears only in trouble says
  nothing in quiet. All 38 frames have the footer now, which also gives 31 of
  them the visible version line they lacked (C-13); on a narrow screen it is
  hidden with the rest of the rail, where there is no free slot for it. The
  nine states stand together in the new frame `components-35`, so frames are
  194.
- **"Match threshold" is out of the rail footer.** It sat in the footer of
  `s-dup-1` and `s-dup-2` as a form field — the one thing that footer is not
  for, and nowhere anyone would look for it. It is a control over the list, so
  it sits in the head under the kind filter over the same list (new
  `tc-acts--stack`, `tc-field--inline`) rather than in a bar of its own — a
  whole row for one field was too much — and the head stopped repeating its
  value in prose.
- **Dates in frames are written one way**, `2026-08-16 09:12` — the form the
  device table and the backup history already used; four dotted ones were left.
  `tc-mono--ink` lost its last user when a date in a banner started taking the
  banner's tone instead of forcing ink, and `dead-rules` said so.
- **The davloom mark is drawn from its guide, and the app rail shows its
  state.** The sprite had a three-thread stroke sketch in a 24 grid, drawn at 13
  and 15 px; the guide (Carrel `docs/visual/davloom-icon.html`, v1.1) has five
  threads, a filled beam and a filled accent cloth, and does not go below 16 px,
  which has its own hinted geometry. Under the icon set's `fill:none` the beam
  and cloth cannot survive, so the mark is built the way the Carrel logo is:
  parts in the sprite, a `<use>` per part, colours from `--tc-ink` and
  `--tc-accent`, which are the guide's colours in both themes. It replaces the
  sketch on the settings rail item and joins the footer of the app rail — as a
  mark, not a button, pulled right so it does not read as one more action:
  absent when davloom is off, present when it runs, and on a danger plate —
  tinted ground, red border — when it does not answer. The mark itself is not
  recoloured: a red cloth at 7×4 px was tried first and could not be read. All three are frame `components-36`, so frames
  are 195.
- **The app rail footer has one row of buttons, the same on every screen.**
  Over 27 frames it had drifted into four shapes: three buttons on ten, two on
  seven, only "Connect account" on three, and on seven a lone plus titled
  "Connect account". The first button's title had been copied as "New
  calendar" onto notes, profile and desktop frames. The row is now: add what
  the rail lists, named by the section the rail is in (address book, calendar,
  list, notebook, folder), connect an account, subscribe to a feed — 17 frames
  changed; a state already drawn on a button (`is-focus` on the profile frames)
  stays.
- **The width ladder had not been run on the rail-footer commit, and it finds
  two things there.** The Match threshold field used `tc-in--chip`, whose fixed
  22 px beat the bar's own heights, so at phone width it had no 44 px tap zone
  (Р-21); it is `tc-in--w52` now and takes its height from the bar like every
  other bar field. And the banner strips in `components-35` were drawn as a
  `tc-shell__rail` outside any shell, which the ladder rightly reads as a rail
  without a shell mode; the strip is the footer itself at rail width
  (`tc-railfoot--strip`).
- **The duplicates frames say what S-16 and S-23 do** (`s-dup-1…6`). Points
  follow the instance weights — email and phone 60, UID 80, name 20, birthday
  10, title and date 70 — and the threshold is 60, with the six weights and
  their instance defaults in a menu under a "Weights" button, as the column chooser is (`s-dup-3`). Only real
  questions are asked (name, organisation and photo of a contact, title and
  place of an event, title of a task or note): no Due, no Tags. The target of
  a merge is chosen on the merge screen ("Keep … instead"), not by a picker on
  the group, and a group with a read-only member cannot be merged and says why.
  Decided groups stand on the same screen after the open ones: linked ("Keep
  these values", "Unlink") and not duplicates ("Offer again"). New frames: an
  event merge that names the reminders gained and the occurrences not carried
  over (`s-dup-4`), a merge that did not end in one record with "Still there"
  (`s-dup-5`), and "Nothing is loaded yet" for an empty cache (`s-dup-6`).
  `tc-merge` lost its last user and is gone. Frames are 206.
- **Mail, publishing and admin settings, as S-26 builds them.** The sheet of a
  link just made — the link, its expiry, "Open the link", "All published
  links" and "Send by letter" — is a frame now (`s-publish-13`), and so is the
  calendar-for-a-range attach sheet (`s-mail-10`): one calendar picked on the
  left, From and To with "Both days are in the file." and "Show" on the right,
  as mail_attach.go builds it. The list holds only the person's own calendars:
  a subscription is not theirs to attach. A whole calendar is not an attach
  kind; it goes by range. The captions say where errors go: a failed attach disables the
  sheet's first action with its reason, a failed send is the first phrase of
  the limits line. Admin settings gain the read-only "Public address" row
  (`CARREL_PUBLIC_URL`) and the test letter as a form of its own, its button
  off with a reason while mail is not configured (`s-admin-3`). The last dotted
  date in the frames is ISO now. Frames are 202.
- **Linked records read as the code builds them.** The block under a card
  lists every other member of the linked group — its stamp, a link to that
  record, and "Phone, email differ", "Holds the same" or "The server did not
  answer; the record stays linked" — instead of "open" and "phone differs"
  (`s-contacts-6`, `s-contacts-7`). Events, tasks and notes have the same block
  now (`s-agenda-2`, `s-tasks-2`, `s-notes-2`), and the task and note lists
  show a linked record as one row with "linked ×2" and "2 sources"
  (`s-tasks-1`, `s-notes-1`).
- **The davloom frames say what plan S-27 and the code do** (`s-davloom-1…3`).
  A proxy device has no rights of its own — it can do what the person can do
  in their connections — so the proxy-wide "Read only / Read and write"
  ceiling, the same switch in the add dialog and the Access column are gone.
  Revoking deletes the record, so there is no "revoked … Delete" row. The tree
  shows the code's levels (`principal/`, `cal/`, `card/`, `files/`) and each
  collection at `c/{account}/{col}/` — connection ID and catalog path, the
  name only beside it. The add dialog ticks collections as well as kinds, as
  the code does; step 2 is the password, the address and the user name with
  "Done", without a code (empty by plan) or setup notes (not in C-6). Five
  rules lost their last user and are gone.
- **The new-device frames say what the server does** (`s-plan-1…5`, task 4 of
  `docs/plan/typecase.md`). An account is included or left out whole — the
  profile and DAVx5 take the account — so the tick sits on the account row,
  titled `label · login`, and its collections are listed in it. The pages no
  longer show a login and a password per server with Copy and Hide: there is
  one "Show password", which opens the unlock window when locked and the
  Connections page inside it. The Apple profile has no password field of its
  own (it is issued inside the unlock window, as `s-plan-3` already drew), its
  link lives five minutes, not ten, and opens once for the person signed in
  to Carrel, and a ready link offers "Download profile" and "Revoke" — nothing
  to copy. Android shows one code per connection, with "Open in DAVx5" on the
  same link and "Code for …" for the others. Two rules lost their last user
  and are gone (`tc-in--t12-w190`, `tc-table--flush`).
- **Mentions, a candidate function, have a section of their own (7.23).**
  Where the timeline shows a person as the subject — attendee, organiser —
  mentions show them as the object: listed in another card (`RELATED`,
  `MEMBER`) or named with `@` in a note, an event or task description, or a
  card's note. Five frames (`mentions-1…5`, frames 200): the inset on the card
  next to the timeline, the screen, the `@` suggestion (the "Related to" menu,
  people only), a mention as it reads — a link to the card — and as the server
  stores it, `@[Name](urn:uuid:…)`, and the empty state that tells "nobody"
  from "not asked". No other frame changes, and the candidate frames carry no
  `data-screen`, so nothing in Carrel checks them yet.
- **Connection labels in the frames are one word** — `nextcloud`, `baikal` —
  because that is what people usually call a connection; `radicale · vps` and
  `sftpgo · vps` stay compound on purpose, as the hard case. The contact row's
  source cell gets its `tc-row__source` class, the anatomy caption of
  `components-12` names the real order in which columns leave, and the rule is
  in `mockup.md`.
- **A list does not mark a duplicate candidate.** `s-agenda-1` showed the
  same duty from two calendars folded into one row with "2 sources" and a
  "duplicate? ×2" badge — but only linked records fold, and the owner decided
  against the hint: a question hanging in sight while one works distracts, and
  a pair found by mistake would distract for nothing; when two records get in
  the way, the way is the Duplicates screen. The duty is two plain rows now,
  each with its calendar, and the duplicates screen no longer says the agenda
  folds it.
- **Six of the frame fixes the server work asked for
  ([docs/plan/typecase.md](docs/plan/typecase.md)).** The origin stamp is one
  form everywhere — connection label, then collection, `nextcloud · home ·
  Контакты`: 168 truncated stamps in list columns, head subtitles and error
  lines, and 27 reversed ones in the "where" pickers, are brought to it. The
  return link says "Back to list" from a card and "Back to the card" from an
  edit form, on every kind. The back rail of the contact and event card and
  edit forms carries every collection of its connection. The quick note is the
  common new-record sheet: "Open the card", no title hint, route
  `/app/{kind}/quick`, and no "where" when there is nowhere to write. And the
  quick-sheet caption no longer says the record is saved before "Save".
- **A check with a hint under it stood on the label line, not on the field
  line.** In a row of labelled fields, a check alone gets an empty label line
  and the 28 px row height, so it lines up with the inputs; the rule matched
  only a check that was the field's only child, so "All day" with its hint
  (`components-10`) floated up to the labels. The rule now covers a check
  followed by a hint. Fields where that same pair stands on its own line
  (`s-profile-10`, `s-admin-3`, `s-backup-2`, `m-desktop-6`) are marked
  `is-own-line`, which is what the mockup already did for such cells;
  computed styles confirm only `components-10` moves.
- **Seven frames had their main area knocked out of the shell grid, by the
  rail-footer commit above.** Where a footer already existed, the rewrite
  replaced it only up to the first `</div>`, leaving the old version line and
  one `</div>` too many. The browser closed the rail early and moved everything
  after it out of the grid (`s-davloom-1`, `s-backup-1`, `s-admin-2`…`-5`,
  `-7`). No check saw it: the compare skipped those frames, and the frame parser
  counts only tags of the frame's own name, so the frame still came out
  balanced. Fixed, and the new `nesting` rule below keeps it from coming back.
- **The backup screen says "backup", not "copy".** The rail item, the page and
  the banner now agree, and so does the table column that used to read "Last
  copy". The key-escrow sentence "No copy of your data key is deposited" is a
  different subject and stands.

- **The screenshot compare moved its base to the current mockup, and the
  exemption table is empty again.** It had drifted to the point of proving
  nothing: of 194 frames, 163 existed in the base commit from 11 September and
  130 of those were exempted by a row, so a forty-seven-minute run was
  comparing 33 frames. The rail-footer change touched 41 and the run went green
  without comparing one of them. The confirmations those rows held are in this
  file and in the history; `allowed-diffs.md` holds what is being compared now,
  and says when to move the base again.
- **A compare with nothing to compare now skips.** With a base set and the
  mockup identical to it, the test used to take one screenshot of every frame
  at both widths and throw them away — twenty-odd minutes. It skips in a second
  and a half instead. That the frames still render is `TestMockupLadder`'s job,
  and that takes ninety seconds for all 194.

### For maintainers (checks)

- **`frame-elements` lets a menu row that holds a field be a `div`**, as it
  already did for a row holding a segment: a control is never nested in a
  button. The weights menu of the duplicates screen is the first such row.
- **New lint rule `nesting`: the whole file must be well nested.** Every
  closing tag closes the element opened last, and nothing is left open. It
  exists because of the seven broken frames above; read strictly, the rest of
  the file was already well nested, so the rule costs nothing. Against the
  broken commit it reports the nine places; against the fix, none.
- **A rename is checked by what the CSS computes, not by screenshots.**
  `TestMockupComputedStyles` takes the mockup at `HEAD` and in the working
  tree, injects each frame at 1351 and 820, and compares every computed
  property and the box of every node. For a change that is not supposed to
  move anything the tree is the same tree, so matching properties mean a
  matching picture — and a difference is named as the node and the property
  instead of a count of pixels. Five minutes against the compare's half hour.
  It found two real mistakes in the rename: a name that collided with the
  existing `.tc-field--full` (`flex:1 1 100%;max-width:760px`), and a rule
  written twice. It also had two faults of its own, both now fixed: it compared
  the properties in the order `getComputedStyle` lists them, and custom
  properties are listed in an order that depends on which rules declare them,
  so adding a rule reported thousands of differences with nothing in them; and
  it measured before the webfonts had loaded.
- **Each `--rN` class was asked whether it does anything.**
  `TestNumberedClassEffect` reads the properties a numbered rule declares,
  takes the class off the node and reads them again; it is skipped unless
  `TC_NUMBERED_EFFECT=1`, so it costs CI nothing. Of 602 classes, 29 declared
  what the node already had and 565 changed the drawing. Among the 324 used
  exactly once the split was 12 against 309 — which is why folding the one-offs
  away is a pile of decisions rather than a mechanical pass, and why it is
  still open.
- **The two long mockup tests have their own CI job and their own timeout
  (Х-6).** `TestMockupLadder` (up to 40 min) and `TestMockupRefactor` (up to
  35) ran inside `go test ./internal/browser` under the default 10-minute
  timeout. The main step now skips them and a `mockup` job runs each with
  `-timeout 60m`, on pushes that touch the mockup, its allowed differences, its
  fonts or the code that reads them — and always when the base commit cannot be
  worked out.
- **A hung screenshot is called a hang, not a difference.** The compare reports
  `browser.IsHang` failures in their own words and gives the frame one more go
  in a fresh browser process before failing it. On a full 165-frame run this
  caught `m-desktop-9 @820` and said so, while every other frame compared clean;
  run on its own that frame passes in thirteen seconds.
- **`instance.stop` kills the browser that outlived its process tree.** A
  browser that re-execs itself leaves the tree `taskkill /T` walks, keeps the
  `typecase-browser-*` profile locked and hangs the next frame. Stop now also
  looks for processes holding that profile — matching browser process names
  only, or the query finds the shell running it (`parseProcessList`).
- **A shallow clone says so.** Both CI jobs fail early on one, and
  `LoadMockupHTML` names `fetch-depth: 0` instead of passing on git's
  "bad object" when it cannot read the compare base.

## typecase/v0.6.0 — 2026-09-21

Geometry comes from `layout` tables locked to the mockup, the remaining mockup
sections are assembled as demo screens, and the mockup itself is rewritten in
the Typecase dictionary so the next release can take its CSS verbatim.
Breaking for consumers: see the `Field`, `ButtonDanger`, row-link and
`tc-fields` items below.

The typecase wordmark and icon work documented below stays documentation-only
until released separately.

### For maintainers (mockup source)

- **The mockup plan is closed; its rules live in
  [docs/visual/mockup.md](docs/visual/mockup.md).** Read it before touching
  `docs/visual/carrel-ui-mockups.html`: what the file holds, the `tc-*`
  dictionary, the closed decisions (Р-3 shell states, Р-4 bar overflow, Р-21
  tap zones and the rest), the width ladder, the lint rules, the before/after
  compare and how to change the mockup. From now on the mockup is edited in
  Carrel and the copy here is replaced whole.
- **What the closing review found is a plan, not a footnote:**
  [docs/plan/1-mockup-tails.md](docs/plan/1-mockup-tails.md). The lint is
  green, but 596 numbered `--rN` classes stand where step 5 asked for named
  variants, the lock deleted dark-theme rules for avatar, thumb and scrim
  because no dark frame shows them, `.demo` blocks are only half under the
  frame rules, and `data-tc` carries empty `{}` rules. Do these before plan 2
  cuts the CSS.
- **The library-from-mockup release is 0.7.0**, not 0.6.0: plans 2 and 3 say
  so now.
- **`mockup-parity` no longer asks for the no-rail + side shell.** No frame
  draws it, so the lock removed the rule as dead and the gate was red since.
  `layout.ShellNoRailHasSide` has no mockup number to match until a frame
  shows that shell.
- **The proof reads the mockup it was proven against.** `internal/proof`
  translates the `m-*` dictionary, so its tests and `go run ./cmd/proof` were
  broken from the rename on. It now takes the mockup at `proof.MockupRev` with
  `git show`; a shallow clone skips its tests. `/compare` shows the same
  zeros as before. Plan 3 still deletes it.
- **CI** checks out full history (the compare base and the proof revision are
  older commits) and now tests `./layout`, `./internal/mockcss/...` and
  `./internal/proof`; they were not in the list, which is how the proof went
  red unnoticed. The two long
  mockup browser tests still run under the default 10-minute timeout there —
  see the tails.
- **Shoot mockup frames one after another, with nothing else using the
  browser.** On closing, the ladder passed and the before/after compare showed
  no pixel difference outside `allowed-diffs.md`, but under load a shot hangs
  (`context deadline exceeded`), and a hung run can leave a headless browser
  with a `typecase-browser-*` profile that makes the next frames hang too.
  Stop such processes, then re-run the failed frames one by one:
  `go test ./internal/browser -timeout 8m -run 'TestMockupRefactor/^s-auth$/^8$'`.
- **Nothing to change in widgets, CSS or screens.** Plan 1 step 8 locks
  the mockup: every live lint rule is on (`tc-*` / `is-*` / `mk-*` in
  frames, a class on each node, no `style=`, no `<script>`, named
  `@container`, tokens only, unique `data-frame`, every frame class
  defined, every `data-tc` rule used in some frame). Leftover `--rN`
  stay as one-off classes. Chrome before/after vs the step 1 base is
  0-diff except the rows in `docs/visual/allowed-diffs.md` (step 7
  caption `min-width:0` on 1351 is recorded; the person looked, no
  visible change). The screenshot harness uses one Edge process per
  side so a second Navigate does not hang CaptureScreenshot.
  `go test ./internal/browser -timeout 40m -run TestMockupLadder`.
  Plan 1 step 7 makes the
  mockup the source of narrow behaviour. Named containers are `tcbox`
  (`tc-root`, `tc-main`, `tc-side`, `tc-list`, overlay, sheet, dialog).
  Shell states (Р-3) live in mockup CSS: `--tc-shell-mode` on
  `.tc-shell__body` is `full` / `rail` (641–1007, 56 px icons) /
  `drawer` (≤640). Drawer descendant rules use
  `@container style(--tc-shell-mode: drawer)` so a 560 px sheet or a
  440 px side panel on a wide root does not pick them up. Tap zones grow to
  44 px in drawer without growing the drawing (Р-21). The page ladder
  (radios above the TOC) sets `.frame` width. New frames `m-phone-8` and
  `m-phone-9` show the icon rail and the sources overlay. Class `is-phone`
  is gone; phone look is width 390. `.tc-top` is not a `tcbox` container:
  otherwise padding made drawer chrome start above 640. In `rail` the
  search keeps equal brand/actions wings. In `drawer` the section rail
  becomes the bottom bar; icon-only header buttons are 38×38.
  `.tc-close` sits in the header row on sheets and side panels. After
  `.tc-push` it does not take a second `margin-left: auto`, so Apply and
  destination picks stay next to the close. In `drawer` the header is
  54 px: 38×38 buttons sit in 8 px chrome, the same as the gap between
  them (Р-21 keeps the drawing 38). `.tc-usermenu` has the
  same 10 px inset as a button, so the name is not flush with the
  previous control. The page ladder sets `.frame` width; `.tc-root` is
  `100%`, so a long frame label cannot make the chrome narrower than the
  card. A header button with an icon and a label (2 min ago) keeps
  `max-content` width; idle search collapses to a 38 px icon (tap opens
  search), not `display:none`. The search field gets the same 44 px tap
  overlay as `.tc-field` (an `input` has no `::after`).
  Chrome ladder:
  `go test ./internal/browser -timeout 40m -run TestMockupLadder`.
  Plan 1 step 6 closes
  mockup slips. `.tc-drop` menu offset is **4 px** (Р-7). A mark in
  `.tc-head__title` is **20 px** (Р-8). A file-row icon is **15 px**, as in
  the table filename (Р-10). Table `.tc-acts` is nowrap (Р-11). `.tc-btn`
  sets `text-transform: none`, `letter-spacing: normal`, `font-weight: 400`
  so a button in a rubric is ordinary type, not small caps (Р-12: Attach,
  server-card actions, and the approved extras Show 3 and New folder).
  Chrome before/after vs the step 1 base is 0-diff except the rows in
  `docs/visual/allowed-diffs.md`.
  Plan 1 step 5 strips `style=`
  from mockup frames. Colours and meter widths become `data-tc-colour` /
  `data-tc-width` (page `[data-tc-*="…"]` rules set the variables; product CSS
  reads `var(--tc-colour)` / `var(--tc-width)`). Layout glue is `tc-line` /
  `tc-stack` / `tc-push` / `tc-pad` / `tc-grow` / `tc-cols--*`. Repeating tweaks
  are block variants; one-offs are `--rN` (Р-13) — leftover classes are not
  selectors for new rules. A table that must share columns after `is-2nd`
  slots hide is `tc-table--fit` (break-word, `td.tc-acts` as a table cell).
  Other tables keep the step 1 wrap (`overflow-wrap: anywhere`). Table
  columns are role classes; SVG `stroke="var(--tc-*)"` becomes
  `tc-ico--accent` and kin. A menu row that hosts a segment is a
  `div.tc-menu__item` (a button cannot wrap buttons). Lint `frame-no-style`.
  `go run ./internal/mockcss/inline` is idempotent. The Р-13 visible list is
  approved. Chrome before/after vs the step 1 base is 0-diff except the
  step 4 rows and the Р-13 frames in `docs/visual/allowed-diffs.md`.
  Plan 1 step 4 puts real
  controls in mockup frames: `button` (btn, seg option, check, menu item,
  pick, close, tick), `input`/`textarea`, `a href`, `kbd`. Search and icon
  fields keep a `tc-q` input (value for painted text). Menus that sat inside
  a button become `.tc-drop` (offset 4 px, Р-7).
  Control reset is the proof `controlCSS` (`.tcx` → `.tc-root`).
  `go run ./internal/mockcss/elements` is idempotent. Lint `frame-elements`
  requires those tags in frames. Chrome before/after vs the step 1 base is
  0-diff except date/time native chrome, extracted pop-menu item icons
  (`s-profile-1`, `s-backup-8`), and ≤4 PNG pixels of input ClearType at
  quarter scale — all in `docs/visual/allowed-diffs.md`.
  Step 3 renames mockup classes to the Typecase dictionary (`tc-*`, states `is-*`) from
  `docs/visual/rename.tsv` via `go run ./internal/mockcss/rename`; a second
  run is a no-op. Desktop window chrome is page classes `mk-*` (not in
  `data-tc`). The lint rejects leftover names inside frames (`.mock`, `m-*`,
  short classes) and `mk-*` outside `m-desktop` frames.
  Step 2 puts component rules in `<style data-tc>` in source order and page
  chrome (`@font-face`, `.wrap`, `.frame`, TOC) in a separate `<style>`.
  Step 1 renamed tokens to `--tc-*`, set `data-tc-theme="dark"` on the dark calendar frame,
  and shipped Public Sans, Literata and Source Code Pro next to the mockup
  (`docs/visual/fonts/`). `--tc-dusty` keeps the mockup value `#6A6351` (Р-9).
  Step 0 numbered frames and added the unique-id lint and the Chrome
  before/after tool; allowed exceptions live in `docs/visual/allowed-diffs.md`.

### For consumers (v0.5.0 → v0.6.0, layout)

- **Do not import `layout`.** Geometry tables are the library's, like `token`.
  Screens keep calling `widget`.
- **`Cols` is the fallback when the catalog has no form.** Three shapes:
  `ColsHalves`, `ColsThirds`, `ColsSidebar` (fluid + 296 px, the same column
  as the document meta). Pass `Items` — one node per track. Empty `Items`
  renders nothing. Widths and gap are not fields: you pick a shape. `Cols`
  does not replace `Shell`, `Doc` or `Fields`. A wrapping tile grid is still
  `Grid`. When a screen needs a local form, it names a type in its own
  package and builds it from `Cols` and primitives.
- **Row variants now in this drop:** `RowAvatarCols`, `RowTickCols`,
  `RowTimeCols`, `RowTickMarkCols`, and `RowFind` as a grid. A grid row must
  live in `List`. `Menu` on a grid variant fails the gates instead of
  rendering nothing. Contacts that used `RowAvatar` (two lines of text) and
  want the mockup columns use `RowAvatarCols` with the cells filled in.
  Search suggestions are `RowLead` (title + subtitle); `RowFind` is the
  `/find` hit list, not the popover. `RowAvatar`, `RowLead`, `RowTick` and
  `RowMeta` stay as flex and remain legal outside `List`.
- **Demo screens `/agenda`, `/tasks`, `/find`.** Week agenda is `RowTimeCols`
  in a `List`, with `TickRow` sources and a `Fields` side; Day/Week/Month/Range
  is `Seg` with Week chosen (no month day grid). Tasks are `RowTickMarkCols`.
  Find hits are `RowFind` in a `List`, not the search popover. No new widget
  types. The Calendar nav item points at `/agenda`.
- **Remaining mockup sections are assembled** (step 9): `/conflict`, settings
  `/account` through `/appearance`, `/mail`, `/admin`, `/login`. No new widget
  types. Where the catalog had no form, the screen named a type in its own
  package and built it from `Cols` (`deviceSetup` halves, `adminWays` thirds).
  Sign-in is Brand + Form, without the shell. The `layout-orphans` gate prints
  a report — catalog widget on ≤1 screen, without `data-tc-`, `role`/`aria-`,
  or a layout number — and does not fail; look at `go test ./internal/gate -v`.
- **Shell, search, rail and side** take their geometry from `layout` (step 4).
  No new widget types. Search shows the 13 px icon as the first node of the
  field. `NavLayoutBottom` matches the mockup tab bar (icon above the label).
- **Form, field, sheet and dialog** take their geometry from `layout` (step 5).
  `Field.Width` is `FieldWidth` (xs/sm/md/lg/xl/full, same six steps as
  `PFieldWidth`, which gains `XL`). `OwnLine` drops the form spacer on a lone
  button. Overlay `Dialog` is 420 px. Sheet target is 200×26 with pick
  chevrons as pseudo-elements, not a shared arrow widget. `/io` wraps Into
  in `Form`. Date/time/number inputs no longer force `10em`.
- **Chip, tag, empty, table, tree, grid and document** take their geometry
  from `layout` (step 6). Chip and tag match `.m-tag` (10.5 px, padding 0 5,
  min-height 16; `is-on` is accent on both). Empty is left-aligned. Message
  has a card ground and a 2 px left bar; alert is a full danger frame.
  Skeleton lines are 9 px. The run meter is 64×3. Table has no wrapper
  frame; header 10 px / 0.12em; cell padding 7 12. Tile grid is
  `repeat(auto-fill, minmax(132px, 1fr))` with gap 14; thumbs are 4/3.
  Tree rows are 26 / gap 7. Queue is 330 px, absolute in `.tc-main`.
  Document layout is `minmax(0,1fr) 296px`; the reading title is 27 / 1.15.
  List rubrics match `.m-group`. Line colours come from `layout` (`--tc-edge`,
  `--tc-hair`, `--tc-rule`).
- **Breaking, now in this drop:** `Field` label sits above the control and
  takes `Width`. `ButtonDanger` is outline; fill is `ButtonDangerSolid`.
  Delete actions in the library (danger zone, danger dialog confirm,
  document meta) default to `ButtonDangerSolid`.
- **`Field.Bare` and the field's accessible name.** A field's `Label` is now
  always the accessible name of its control; before, the control carried
  `aria-label` from the machine `Name`, which overrode the visible label.
  For a field without a visible label, keep `Label` and set `Bare: true`: the
  label leaves the screen, the name stays. An empty `Label` still renders
  bare, but then the control is named by `Name` — do not write fields that
  way.
- **Hints and errors are tied to their control.** `Field` and `PField`
  (including a `Check` or `Pick` inside) put `aria-describedby` on the
  control: hint first, then error. Their ids now derive from the control id
  (`tc-field-N-hint`), not from `Name` (`name-hint`), so two fields with one
  `Name` no longer share an id. If your script looked up `#<name>-hint`, look
  up the control's `aria-describedby` instead. No new input fields.
- **A row link is the row's main text, not the row.** `Row` and `TickRow`
  with `Href` used to render the whole row as `<a>`, with the tick, the tags
  and the menu inside it: invalid HTML, and a press on the tick followed the
  link. Now the root is a `div`; the title (`name` for `RowMeta` and
  `RowTickCols`, the label for `TickRow`) is the link, stretched over the row
  with `::after`, so the whole row still opens it, and the tick, tags and
  menu sit above and work on their own. Nothing changes in your code. If your
  CSS or scripts targeted `a.tc-row` or `.tc-tickrow__link` as a wrapper,
  target `.tc-row__link` / `.tc-tickrow__link` on the text instead.
- **`tc-fields` label column is 104 px.** It is no longer `8em` and no longer
  grows with the longest label. Pair lists that assumed a flexible label
  column will wrap or clip sooner.
- **Geometry is locked to the mockup.** `mockup-parity` compares every pair
  in the layout plan to `.m-*`. The look probe runs assembled screens at
  1600, 1280, 900 and 480; expected sizes come from `layout` for the measured
  container, not a third copy of the numbers. A grid row's slot count must
  match the computed tracks. Nothing to change in your widgets if you already
  compose from the catalog.
- **The fence now covers screens too.** The demo screens fail the gates if
  they import `layout` or carry `style=`, inline handlers or `<script>` in
  their markup — the same rules an application screen follows.

### For consumers (wordmark and icon preview)

- **Nothing to do for the wordmark.** No module code, token, markup or behaviour
  changed for the wordmark; it is not part of this library release.
- **The wordmark belongs to the library's own pages** — the widget catalog, the
  README, a release page — and nowhere else. It marks the library, not an
  application built on it, so it does not go into a consumer's interface, and it
  never stands next to the marks of the application that depends on it.
- **Do not set the word in Literata and call it the mark.** The nests, the air
  inside them and the empty nest are measured, not typed; `typecase-logotype.svg`
  is the only file that carries them.
- **Below 164 px wide, do not use it.** That is where the hairline of "e" still
  holds 1.5 px. There is no hinted small version of the word; in square formats
  the icon carries the library instead.
- **Nothing to do for the icon either.** No module code, token, markup or
  behaviour changed for it; it is not part of this library release.
- **The icon is the square-format mark** — the catalog, the README, a section
  head. It does not replace the Carrel logo and it is not the project's favicon:
  Typecase ships separately, but the project has one logo.
- **Do not draw the letter.** It is the Literata glyph of the same instance as
  the wordmark, scaled to a quarter. Only the cap height is assigned; the width
  and the air follow from the glyph.
- **The box is not a backdrop.** Nests never sit straight on paper or on an
  application's background: card on paper measures 1.18:1, which is no visible
  edge at all. On the ink box the same nest measures 12.47:1.
- **Below 24 px the accent is gone.** The 16 px drawing collapses the case to a
  single nest with the letter — no empty nest, no accent. Expect the mark to say
  "a sort in a case", not "there is room in the case".

### Included

- `docs/visual/typecase-logotype.svg`: the master. Eight letters, each in its own
  nest, and one empty nest as the accent. Three fills, no strokes, no embedded
  text and no metadata, 8.1 KB.
- `docs/visual/typecase-logotype.html`: the construction reference — the single
  source of truth for the mark, in the genre of the Carrel references next to it.
  It records what was chosen and why, including the five rejected accent
  placements and the two rejected nest geometries.
- The instance is `opsz 12, wght 400`, and the reference shows the measurement
  behind it: across the optical size axis the hairline of "e" varies from 5.89 to
  1.34 units, so at `opsz 72` the word would need 564 px of width to keep that
  stroke at 1.5 px, against 164 px here. It is also the instance that sets the
  library's own text.
- The grid is derived, not designed: a nest is the letter's ink plus the
  instance's mean side bearing (3.24) on all four sides, which makes the nests
  sum to the font's own advances exactly (420.42 against 420.42). The partition
  between nests and the corner radius are one module — the stem of "p", 8.77.
- The accent is an empty nest after the word, as wide as the average nest
  (52.55): in a case, an empty nest means the sort is already set in the line and
  its place is still there.
- `docs/visual/tools/`: `typecase_logotype.py` builds the mark from the font,
  `build_typecase_logotype.py` releases the master and the reference together,
  `checks.py` measures contours against the font, containment, gaps, the sums and
  the size limit. Every number printed in the reference comes from that
  measurement, so the document cannot drift away from the file it describes.
- `docs/visual/typecase-mark.svg`: the icon master. An ink box, nine nests, one
  letter and one empty nest as the accent. Four fills in three colours — the box
  and the letter share one — no strokes, no embedded text and no metadata,
  1.4 KB.
- `docs/visual/typecase-icon.html`: the construction reference for the icon, in
  the same genre. It records what the 3×3 sketches got wrong — the 64 field, the
  1.5 stroke that puts every edge at 3.25 and 20.75, the letter set in Georgia,
  the palette that belongs to no theme, and the accent painted on the occupied
  nest instead of the empty one.
- The grid is assigned and the letter is not: 6 + 24 + 6 + 24 + 6 + 24 + 6 = 96,
  where 6 is both the outer field and the partition, and the radius is half a
  module. The only assigned number of the letter is the cap height — 16.5, a
  quarter of the wordmark's 66; the width (15.42) and the air (4.29 by 3.75)
  follow from the glyph, whose contour matches the font's to 0.0065 %.
- The accent rule is the wordmark's: the empty nest carries it, the occupied nest
  never does. In the icon the letter sits in the left middle nest and the empty
  one in the right middle — the line is started and there is still room.
- Colour roles differ from the wordmark, and the reference measures why: the box
  takes the darkest token of the theme, the nests the lightest, the letter the
  colour of the box. That is 12.47:1 for a nest against the box in light and
  13.33:1 in dark, against 1.18:1 for a nest lying on paper.
- `docs/visual/tools/`: `typecase_icon.py` builds the icon from the same font
  instance as the wordmark, `build_typecase_icon.py` releases the master and the
  reference together, `icon_checks.py` measures the glyph against the font, the
  grid, the halftones and the contrasts, and `raster.py` is the rasterizer it
  uses — ready-made ones fill a polygon together with both borders, so an edge
  lying exactly on a pixel never comes out clean.

### Not yet included

- Overflow of `tc-top` / `tc-seg` at 375 px — not the layout plan.
- The mockup tails: [docs/plan/1-mockup-tails.md](docs/plan/1-mockup-tails.md).
- Raster frames of the icon. The 16 px and 32 px drawings exist as constructions;
  `.ico` and PNG are not built from them.
- The lock-up of the icon and the wordmark. Both references exist; the spacing
  and the alignment between the two are described in neither.
- The nest contrast inside the wordmark itself. On paper the nests measure
  1.18:1 in light and 1.26:1 in dark — they disappear and only the word remains.
  The icon works around it with the box; the wordmark has to be fixed on its own
  terms, in its own reference.
- The superseded sketches. `docs/visual/typecase-logo-v2.svg`,
  `typecase-logo-bg.svg` and `typecase-wordmark-tiles.svg` are still on disk and
  no longer describe anything.
- A Cyrillic wordmark. There is none, and Literata set next to the mark stays
  text.

## typecase/v0.5.0 — 2026-09-10

Standalone repository layout: the module root is the repository root.

### For consumers (v0.4.0 → v0.5.0)

- **Import path unchanged:** `carrel.dev/typecase`. Acceptance packages
  (`catalog`, `browser`, `gate`, `screens`) live under `internal/`; import them
  only from module tests and demos, not from application code.
- **Run demos from the repository root** with `scripts/catalog.bat` and
  `scripts/screens.bat` (replacing `scripts/typecase-screens.bat` and paths
  that assumed a monorepo parent).
- **The icon reference names tokens as `css/00-tokens.css`**, not a monorepo-relative path.

### Included

- Repository root matches the Go module layout; `internal/` holds acceptance
  packages; docs, scripts, and CI ship with the module.
- `scripts/catalog.bat` and `scripts/screens.bat` for local demos.
- Layout plan ([docs/plan/typecase-layout.md](docs/plan/typecase-layout.md))
  aims at composing screens, not only at a single geometry source.

### Not yet included

- Calendar, tasks, find screens and remaining geometry — see Unreleased
  ([docs/plan/typecase-layout.md](docs/plan/typecase-layout.md)).

## typecase/v0.4.0 — 2026-09-09

Assembled screens match mockup drawing at 1280×900. Notes frames 7.4 close with
`tc-pfield`, `tc-pick`, and `tc-docedit`. Palette gains `--tc-edge`, `--tc-hair`,
and `--tc-dusty`; icons render through the `tc-i-*` sprite. Remaining mockup
geometry is the `layout` package — [docs/plan/typecase-layout.md](docs/plan/typecase-layout.md).

### For consumers (v0.3.0 → v0.4.0)

- **`tc-pfield` / `tc-pfieldrow` for side panels and focus-edit frames.** Stacked
  label-on-top fields (`PField`, `PFieldWidth`) and wrapped rows (`PFieldRow`).
  `SideBlock` accepts `*PField` and `*PFieldRow`; form `tc-field` geometry is
  unchanged.
- **`tc-pick` for relation fields.** `InputPick` on `PField` with `Pick{ Chosen,
  Query, Options, Manual, Open }` — removable chips, a real query input, candidate
  list, manual UID entry as the last row. `tc.js` filters the list and adds chips.
  Localise the chip button with `Labels.Remove` (default `"Remove"`).
- **`DocEdit` for focus editing.** `Doc{ Edit: &DocEdit{ Title, Rows, Text } }`
  replaces `Prose` in editing mode: title input, property rows, textarea.
- **`DocBar.Sub` + `DocBar.Mark` for reading head.** Source line and actions
  render as one frame (`tc-docbar--frame`).
- **`DocMeta` extensions.** `MetaSection.Action` (e.g. Attach on Attachments) and
  `DocMeta.Danger` at the foot of the meta column. `Rubric.Action` for the same
  pattern in `tc-sec`.
- **Redefine the palette if you override the bundle.** Three new required slots:
  `--tc-edge`, `--tc-hair`, and `--tc-dusty` (light/dark values ship in the bundle).
  `token.Validate` fails without them.
- **`Icon` is a sprite id (`tc-i-*`), not a glyph or `#+` href.** Pass ids such as
  `tc-i-plus`; empty string omits the icon node. Nav icons render through the sprite too.
- **`Brand.SVG` replaces `Mark`/`Word` when set.** Inline `<svg>…</svg>` markup is
  inserted as HTML; screens use the Carrel logo from `screens/carrel-logo.svg`.
- **`Top.Refresh`** is an optional poll action between Create and User (`Label` is
  the time caption, `Icon` is `tc-i-refresh`). Zero value hides the button.
- **`Side.Size`**: `SideSizeWide` (`"wide"`) sets 440 px (`tc-side--wide`); default
  stays 400 px. **`Side.Sub`** adds a panel subtitle under the title.
- **`Menu.Kind`** styles the create control (`ButtonPrimary` + `tc-i-quill` on screens).
- **`TickRow.Current`** marks the active source row (`is-current`, `--tc-card` background).
- **Unlabeled `Field`** (`Label == ""`) omits the label node and the 10em column
  (`tc-field--bare`): control and hint share one left edge.
- **`tc-tickrow .tc-mark`** is the 3×14 px collection bar, not `--tc-row-height`.
  Tickrow padding is 12/14 px. Empty `DocMeta` renders nothing.
- **Excerpt `Row.Badges` live in the tags cell**, not after the grid. The excerpt
  row is a six-column grid that includes the 3 px mark.
- **`Mark.Colours`** draws equal horizontal stripes when more than one colour is given.
- **Rank2 bar actions render as buttons first.** `tc.js` stashes them under «More» only
  when the bar runs short; at 1280 px Import/Export stay visible on `/notes`.
- **Assembled screens match mockup drawing** at 1280×900. The browser probe
  (`screens_look_test.go`) checks nav 29 px and `--tc-card` current, tickrow 26,
  `tc-main` on `--tc-bg`, list without border, sprite ids, `/notes` side 440 px,
  Rank2 buttons, head title 22 px / line-height 1.15, and every `/states` mock.
  Demo chrome lives in `.screens-chrome`; the product shell inside
  `.screens-stage` uses `min-height: 100%`, not `100vh` with page chrome.
  Look probe also checks: quick-sheet `tc-field--bare` and textarea/hint
  alignment, sheet textarea min-height 132 px, tickrow mark 14 px and no
  overlap with rail section titles. Search in the top bar is 560 px; the
  shell rail is a flex column so chips and foot stay in the 1280×900 frame.

### Included

- `tc-pfield`, `tc-pfieldrow`, `tc-pick`, `tc-docedit`; `DocEdit`, `DocBar.Sub`,
  `DocMeta.Danger`, `MetaSection.Action`, `Rubric.Action`, `SideBlock` pfield
  kinds; notes frames 7.4 / read / edit assembled (`Э-7`…`Э-10`).
- `screens-compose` locks panel fields on `/notes`, `tc-docbar--frame` on
  `/notes-read`, and `tc-docedit` / `tc-pick` on `/notes-edit`.
- Palette slots `--tc-edge`, `--tc-hair`, `--tc-dusty`; mockup-aligned shell CSS (nav
  29 px, tickrow 26 px, transparent `tc-main`, list hairline, card selection).
- Expanded `icons/sprite.svg`; sprite + logo defs embedded in screens and catalog pages.
- `screens/carrel-logo.svg`; `Brand.SVG` on demo screens.
- `Top.Refresh`, `Side.Size` / `Side.Sub`, `Menu.Kind`, `TickRow.Current`,
  `Mark.Colours`, icon-only railfoot buttons, `DocBar.RightNodes`.
- `gate/geometry`: nav 29, tickrow 26, tickrow pad 12/14, tickrow mark 14,
  sheet textarea min-height 132, list `border-width` 0, head title 22 px.
- `browser/screens_look_test.go`: per-screen look probe at 1280×900, including
  `/states` frames, current-nav background, main surface, head title measure,
  `/notes-quick` bare field, and rail mark height.
- Demo `screens-chrome` + stage layout in `screens.css` / `PageHTML`.

### Fixed

- Rank2 Import/Export were only reachable through the bar overflow menu at desktop width.
- Rail tag chips rendered above sections instead of after them.
- Demo page chrome was counted in `.tc-shell { min-height: 100vh }`.
- Head title used `line-height: 1.5` instead of 1.15; ink overflow is absorbed
  by `padding-block` so the В-1 check does not return.
- Notes reading frame omitted the panel-hide control (`tc-i-panel`).
- Panel date, time and select did not fill their `tc-pfield` cells: native
  date/time stayed 10em, and a class rewrite broke the select wrap.
- `tc-pick` query was a span, so the field could not be typed.
- Focus-edit padded twice (`.tc-doc__layout` and `.tc-docedit`).
- Rubric `Action` inherited uppercase from `tc-rubric`.
- Unlabeled sheet fields reserved a 10em label column, so the textarea sat to
  the right of its hint.
- Rail collection bars used list-row height (38 px) and painted through
  section titles and neighbouring ticks.
- Excerpt `linked` badges sat outside the six-column grid.
- Empty `DocMeta` still emitted a 16em aside beside the edit textarea.
- `.tc-rail { height: 100% }` overflowed chips and the rail foot below the
  900 px frame; top search `flex: 1` stretched across the bar.

### Not yet included

- Calendar, tasks, find screens and remaining geometry — closed in Unreleased
  ([docs/plan/typecase-layout.md](docs/plan/typecase-layout.md)).

## typecase/v0.3.0 — 2026-09-07

Screens assembled from widgets. Geometry matches the mockups; notes frames no
longer bypass `tc-main`, `tc-bar`, or `tc-side`; `tc-seg` and `tc-top` collapse
when width runs short. Remaining demo pages still wrapped the shell by hand;
that leftover is closed in Unreleased. The narrow rail's open control is also
closed in Unreleased.

### For consumers (v0.2.0 → v0.3.0)

- **`Node` and `[]Node` for widget composition.** `Main.Body`, `Card.Body`,
  `Sec.Body`, `Bar.Left`, and `Bar.Right` now take `[]Node` (any value with
  `Render() string`), not `Prose` or `[]Action`. Use
  `widget.Paragraph("text")` as a `Node` inside the slice; for bar actions
  use `widget.ActionsAsNodes([]widget.Action{...})` or
  `[]widget.Node{widget.Action{Label: "Save"}}`.
- **`Sheet.Body` is `[]Node`.** Pass `[]widget.Node{widget.Paragraph("…")}` or
  fields and other widgets; a bare `Prose` value no longer compiles.
- **`RailSection.Body` for non-tickrow rail content.** When `Rows` is empty,
  `Body` renders instead (for example a neighbour `tc-tree`); do not set both.
- **`Top.Narrow` and `tc-top--narrow` are gone.** Narrow layout is computed in
  `tc.js` from width: search shrinks first, then the right group goes under
  `labels.More`. Do not pass window width as widget input.
- **Redefine tokens if you override the bundle.** `--tc-row-height` is now
  `38px`, `--tc-font-size` is `13px`, and three new required tokens appear:
  `--tc-page-gutter` (`20px`), `--tc-page-gutter-narrow` (`14px`), and
  `--tc-radius-sm` (`2px`). `token.Validate` fails without them. Defaults ship
  in the bundle; nothing to do if you use `typecase.CSS()` as is.
- **Default button height is bar (26 px), not form (28 px).** Bare `Button{}`
  and `Action{}` now render `tc-btn--bar`. Form footers still get
  `tc-btn--form` automatically; elsewhere pass `Size: ButtonSizeForm` when the
  control sits in a form.
- **Page gutter replaces ad-hoc horizontal padding.** Panels, rows and headers
  read `var(--tc-page-gutter)`; narrow shells swap it for
  `var(--tc-page-gutter-narrow)` via `.tc-shell--narrow`.
- **Run the screens demo beside the catalog.** `scripts/typecase-screens.bat`
  serves `http://localhost:8081/` — assembled mockup pages, not widget
  cells. Routes: `/notes`, `/notes-read`, `/notes-edit`, `/notes-quick`,
  `/files`, `/contacts`, `/io`, `/duplicates`, `/states`. Append `?theme=dark`
  for a dark rerender before first paint. `screens.css` holds page chrome
  only; it must not contain `.tc-` selectors.

### Included

- `widget.Node` and `[]Node` on `Main`, `Card`, `Sec`, and `Bar` shelves; gate
  `screens-compose` on `/notes`.
- `SideBlock` union (`*Sec`, `*Fields`, `*Field`); `/notes` side panel through
  `Side.Render()`.
- Empty `RailSection.Title` omits `tc-railsec__title`.
- `RowCells.Date` and `Source` on `tc-row--excerpt`; excerpt is one line with
  ellipsis.
- Narrow layout (L-4.4): `tc-seg` collapses to `tc-menu` when width runs short;
  `tc-top` shrinks search first and stashes the right group under `labels.More`.
- Notes demo frames: routes `notes-read`, `notes-edit`, and `notes-quick`; list
  frame drops inline `tc-doc` preview; read frame carries neighbour `tc-tree`
  in the rail; `RailSection.Body` and `Sheet.Body []Node`.
- Mockup-aligned geometry: row height 38, row gap 12, rail tickrow padding 16,
  top bar 52, default button 26, second radius 2 px on check/badge/mark.
- Gates `geometry`, `screens-css`, `screens-coverage`, and `screens-compose`.
- Package `typecase/screens`, binary `cmd/screens`, routes for s-notes,
  s-files, s-contacts, s-io, s-dup and m-states mockup frames.
- Browser check `screens_test.go`: each screen at 1280×900 and 375×812, light
  and dark (`?theme=dark`), including page-level horizontal overflow at 375 px.
- Demo chrome follows the palette tokens, so dark theme does not leave the
  screens nav in light ink. Theme links: `?theme=light` / `?theme=dark`.

### Fixed

- Screens browser probe: SVG `className` is not a string; `color-mix`
  backgrounds come back as `color(srgb …)` and now parse for contrast.
- Completeness: `tc-btn__icon` had no stylesheet rule.
- `tc.js` treated every `data-tc-seg` node as a group, including the option
  buttons; layout and keyboard now bind to `.tc-seg`.
- `tc-bar` restore after collapse: left-shelf items were inserted before the
  overflow menu, which lives on the right shelf, so a second layout pass threw.

### Not yet included

- Narrow rail open control and hand-wrapped shells: closed in Unreleased.
  The 375 px `OutOfBounds` skip of `.tc-main` is also closed there.

## typecase/v0.2.0 — 2026-09-05

The second visual audit pass, closed (`В-12`…`В-27`). The palette gains an
eighth slot, small controls stop drawing their tap area, headings stop
inheriting the browser's scale, and sheets and dialogs are finally positioned
by the library.

### For consumers (v0.1.1 → v0.2.0)

- **Define `--tc-card` if you override the palette.** The eighth palette slot: a
  quiet surface tone one step off the page ground. `token.Validate` and
  `token.ValidateWidget` now require it, so a consumer stylesheet that redefines
  the palette without it fails validation. Defaults ship in the bundle:
  `#e8dfcb` light, `#242e27` dark. Nothing to do if you use the bundle as is.
- **Mark the root with `data-tc-root`, not the `tc-root` class.** Base
  typography, colour and `box-sizing` now hang on the attribute the markup
  already carries, so the library no longer needs a script to look like itself.
  `tc.js` stops adding the class; if your own CSS targeted `.tc-root`, retarget
  it at `[data-tc-root]`.
- **Small controls are drawn at their real size again.** `tc-tick` and the
  tick in `tc-tickrow` are 13 px, `tc-close` is 24 px. The 44 px tap area comes
  from a transparent pseudo-element below 720 px, so layouts that reserved
  44 px of space for these marks will now see them take 13 or 24. `tc-check`
  no longer forces the touch row height at every width.
- **Headings inherit the base size unless a widget names one.** If you styled
  around `h1` being 2em inside `tc-head`, drop those overrides: the title is
  `1.6em`, the sheet and dialog titles are body size in bold.
- **`tc-usermenu` markup changed.** The popover is a `div` wrapping the header
  and a `ul[role=menu]`; `aria-controls` points at the list. CSS or scripts
  that assumed `ul.tc-usermenu__popover` need updating.
- **`Badge` with an empty label renders nothing** instead of an empty chip.
- **`tc-seg` changed appearance.** The chosen part is `--tc-card` with ink text
  and `font-weight: 600`; the others are muted with no fill. A solid accent
  marks the primary action and the "done" state, and a view switch does not
  take that rank.
- **The side column's width moved to `.tc-side`.** `.tc-shell__side` is
  `width: auto`, so a closed panel collapses its column instead of holding
  400 px open.
- **Run `go test ./...`** from the module root after upgrading. The browser
  gates need Chrome or Edge (`CHROME_PATH` overrides detection).

### Fixed

- `tc-tick`, `tc-tickrow__tick`, `tc-close`, `tc-check`: the 44 px tap area was
  written as `min-width`/`min-height` on the mark itself, so it was painted at
  every width instead of being reserved for the narrow layout (`В-14`).
- `tc-tickrow` in the narrow layout spaces the tick and the label by `1.2em`, so
  the tick's tap area does not swallow taps meant for the label.
- Heading sizes: no `h1`–`h4` inside the library inherits the browser scale any
  more. `tc-head__title` is `1.6em/1.5` serif, `tc-sheet__title` and
  `tc-dialog__title` are body size in bold, `tc-empty__title` is `1.1em`, and
  the document heading scale is stated in `09-document.css` (`В-15`).
- `button`, `input`, `select`, `textarea` take `font: inherit`; `code`, `kbd`,
  `samp` and `pre` take the declared monospace. Nineteen controls were drawn
  with the system face (`В-16`).
- `tc-sheet` and `tc-dialog` are centred by their overlay layer, 44 px from the
  top, at `560px`/`92%` and `24em`/`90vw`. The layer is transparent to pointer
  events so the scrim still closes on a click outside (`В-17`).
- `tc-sheet__head`: the title leads and the target caption is pushed right;
  `width: 100%` on the caption used to squeeze it against the title (`В-18`).
- `Badge.Render` with no label renders nothing; `tc-usermenu` no longer emits
  the empty badge that its own `inv-badge-has-text` fixture describes, and the
  popover header no longer sits inside the `ul` (`В-19`).
- `.tc-main` fills its column, like `.tc-rail` and `.tc-side` (`В-20`).
- A closed side panel collapses its column (`В-21`).
- `tc-rail` draws no chip strip when there are no chips (`В-23`).
- Base styles hang on `[data-tc-root]`, so the library looks the same without
  `tc.js`; the static catalog page used to fall back to 16 px Times (`В-24`).
- `.tc-mark` has no palette default: without a supplied colour it is not
  painted, rather than painted with the theme accent (`В-25`).
- `tc-seg`: chosen part follows the mockups, carries a non-colour cue, and takes
  `flex: none`/`white-space: nowrap` so a part is never squeezed below its
  label (`В-26`).
- `tc.js` moves `is-chosen` together with `aria-pressed` when a segment is
  clicked; the visible and the announced selection used to diverge (`В-27`).

### Included

- `catalog.Mux`: the catalog site — pages, styles, scripts, fonts, icons and a
  sample image — built once and used by both `cmd/catalog` and the browser
  checks. Unknown paths answer 404 instead of the page, which is what made font
  faces fail to parse and images fail to decode (`В-12`).
- `--tc-card` in `css/00-tokens.css`, in `token.Palette`, and in the L-5.1
  palette list in `docs/spec/typecase.md`.
- `browser.TapTarget` measures the hit area with `elementFromPoint` at four
  points 21 px from the centre, across five controls, and separately checks the
  drawn size at 1280 px.
- `gate.compositeInvariants`: per-occurrence markup invariants read every
  catalog entry the widget is drawn in, not only the entry named after it.
- Browser checks for catalog assets, typography, overlay geometry, shell
  geometry, the no-script base, and segment selection agreement.
- Catalog fixtures carry a build caption; `tc-railfoot — missing` shows the
  unassembled-input state on purpose.

### Not yet included

- `tc-seg` does not collapse into a menu when width runs short (L-4.4).

## typecase/v0.1.1 — 2026-09-05

Visual audit fixes from the catalog browser review (findings В-1…В-7) and
follow-up gate/browser checks (В-8…В-10).

### For consumers (v0.1.0 → v0.1.1)

- **No breaking API changes.** CSS, markup, and gate checks only. Run `go test ./...`
  from the module root after upgrading.

### Fixed

- `.tc-btn` and ellipsis labels: `line-height` raised so descenders are not clipped
  inside `overflow: hidden` (`В-1`).
- `gate.CheckCompleteness` now requires every `tc-*__*` class in catalog markup to
  have a stylesheet rule or an explicit hook (`В-6`). Bare variant modifiers
  (`tc-row--lead` and similar) stay outside gate scope; see `CatalogClassHooks`
  (`В-8`).
- `.tc-row__text`: column layout so title and subtitle stack (`В-2`).
- `tc-tree`: children moved out of the node flex row (`В-3`).
- `tc-row--avatar` fallback: no longer puts subtitle in `<img src>`; initials use
  the first letter of the title (`В-4`).
- `tc-queue__row--*`: distinct styles per state in the catalog (`В-5`).
- Catalog page: containing blocks for menus and search popovers; browser check
  covers `position: absolute` as well as `fixed` (`В-7`).

### Included

- Browser checks: ink overflow (`fontBoundingBox` vs line box), row text stacking,
  tree label width, queue state fingerprints (colour, weight, decoration,
  opacity), popover containment.
- `gate.CatalogClassHooks` for element classes intentionally without CSS.
- Stylesheet rules for catalog element classes that were markup-only.
- `TestBite_completeness_elementClass` runs through `gate.Run` (`В-9`).
- `TestBite_catalogQueueStates`: removing the idle rule fails the queue check (`В-10`).

## typecase/v0.1.0 — 2026-09-05

First tagged release after the pre-tag reviews. Upgrade in one step from
`v0.0.5` and run `go test ./...` from the module root before deploying.

### For consumers (v0.0.5 → v0.1.0)

- **Body fields are `Prose`, not HTML strings.** `Overlay.Body`, `Sheet.Body`,
  `Dialog.Body`, `Doc.Body`, `Main.Body`, `Card.Body`, and `Sec.Body` now take
  `widget.Prose`. Use `widget.Paragraph("text")` for a single paragraph, or build
  `Prose{Blocks: ...}` for richer content. This is a breaking change from
  v0.0.5, where those fields were unescaped `string` markup.
- **Disabled controls use `Reason`, not `Enabled`.** `Button`, `Action`, `Check`,
  and `Tick` no longer have an `Enabled` field. Set `Reason` to a non-empty
  string to disable; leave `Reason` empty for an enabled control.
- **IDs are generated inside widgets.** `Menu`, `Sheet`, `Dialog`, `Side`,
  `Field`, `Input`, and `Check` assign unique `id` values automatically. Pass
  `ID` only when you need a stable reference for tests or `aria-controls`.
- **`RowMeta` replaces `RowFile`.** The row variant constant is now `RowMeta`
  (`"meta"`); CSS class is `tc-row--meta`. Update any code that used
  `RowFile` / `tc-row--file`.
- **CSS-bound fields use named enum types.** `Brand.Size`, `Avatar.Size`,
  `Skel.Width`, `Rail.Layout`, `Nav.Layout`, `Thumb.Size`, `ProseBlock.Kind`,
  `Step.Status`, `RunItem.State`, and `RailSection.State` are no longer plain
  `string`. `Mark.Colour` and `Avatar.Colour` use `MarkColour`.
- **Localise UI strings with `widget.SetLabels`.** Call before rendering to
  replace defaults such as "Close", "Cancel", "Search", and "More". Use
  `widget.SetStepOfFunc` for step progress text.
- **Run the full check suite** with `go test ./...` from the module root. L-8
  gates live in `gate/`; L-9.2 browser checks in `browser/` need Chrome or
  Edge (`CHROME_PATH` overrides detection).

### Included

- `gate.CheckDocInput`: L-3.2 «Вход» field lists must match exported struct fields.
- `gate.CheckDomainNaming`: const names, enum string values, and L-3.2 «Состояния»
  slugs are checked without the catalog whitelist.
- `gate.CheckInvariantCoverage`: every invariant clause has a `gate:inv-*` markup
  check; `todo:` bindings are rejected; 72 bite fixtures under
  `gate/testdata/inv/`.
- `gate.CheckSpecL4`: implemented widgets must match the L-4 registry in
  `docs/spec/typecase.md`.
- `gate.CheckGofmt` and `gate.CheckVet` run from `gate.Run` via `go test ./...`.
- `gate.CheckOwnership`: scans git repository root, skips the typecase module by path.
- `tc.js`: focus trap lifecycle (`attachTrap` / `releaseOverlayTrap`); no
  autofocus on pre-open overlays at load.
- `gate.CheckUniqueIDs`: catalog HTML must not repeat `id` attributes.
- `gate.CheckA11y`: each button and field in catalog HTML is checked individually.
- `widget.Paragraph`, `widget.SetLabels`, `widget.SetStepOfFunc`, internal
  `uniqueID` for collision-free markup.
- `gate.CheckVerifiable` and bite tests that call `gate.Run` for every gate.
- `tc-table`, `tc-grid`, `tc-queue`, and `tc-runbar`: behaviour moved to
  «Чего нет» (consumer-owned); orphan `data-tc-*` hooks removed.
- `tc-diff`: arrow keys documented as native radio-group behaviour.
- `tc.js`: menu keyboard navigation (arrows, Home/End, first letter), seg
  arrows, tree toggle/selection, search clear, rail drawer, bar collapse,
  side-panel close, crop controls, list arrows, range validation, tickrow toggle.
- `tc-tree` disabled nodes: `disabled` button and `pointer-events: none`.
- `tc-row` avatar/tick variants read `RowCells.Avatar` and `RowCells.Done`.
- `tc-stat` and `tc-stats`: pre-formatted `Stat.Value` string; equal-width row
  with `flex-wrap: nowrap`.
- Catalog page styles served as `/catalog.css` (strict `style-src` on the demo).
- `Column.Kind` defaults to `ColText` when empty in table headers.

### Breaking changes from v0.0.5

- `Body string` → `Body Prose` on overlay, document, shell, and utility widgets.
- `Button.Enabled` / `Action.Enabled` removed; use `Reason`.
- `RowFile` / `tc-row--file` renamed to `RowMeta` / `tc-row--meta`.
- `RowCells` gained `Done` and `Tags`; avatar variant prefers `Avatar` field.

## typecase/v0.0.5 — 2026-09-04

Browser checks that actually run, and a readable catalog.

### For consumers

- The L-9.2 checks now run on Microsoft Edge as well as Chrome. The harness
  starts the browser itself and connects over the DevTools endpoint published
  in the profile directory, instead of relying on the "DevTools listening" line
  on stderr — a browser built as a Windows GUI subsystem binary writes nothing
  to an inherited pipe, so every check previously failed with `chrome failed to
  start` on a machine that has only Edge. `CHROME_PATH` still overrides
  detection.
- One browser process now serves every check in a run; viewports come from
  device metrics rather than a window-size flag.
- The catalog page loads `/theme.js` in `<head>`. The route was served but
  never linked, so the catalog rendered without a resolved theme. Both pages
  carry it: the theme boot script must precede first paint (L-5.3) and is not
  widget behavior. `/static` still withholds `tc.js`.
- The catalog gives each cell holding a scrim, overlay, or rail drawer its own
  containing block. Those widgets are fixed to the viewport by design, so on a
  page that shows every state at once a single open scrim covered the whole
  catalog and hid all 139 entries.

### Included

- `browser`: own launcher (`--remote-debugging-port=0` plus `DevToolsActivePort`)
  behind `chromedp.NewRemoteAllocator`; `chromedp.EmulateViewport` per check.
- `browser/catalog_test.go`: measures that no viewport-fixed widget escapes its
  catalog cell and that the theme attribute is set on load.

### Fixed

- Settle waits in the browser checks were untyped constants passed to
  `chromedp.Sleep`, so `Sleep(80)` waited 80 nanoseconds instead of 80
  milliseconds; menu-open and theme-switch checks measured before the page had
  settled.
- `browser/fixtures.go` was not `gofmt`-clean, failing the first-commit gate.

### Not yet included

- Opus API review.

## typecase/v0.0.4 — 2026-09-04

L-9.2 browser checks (chromedp) and menu popover positioning.

### For consumers

- Run `go test ./...` from the module root to execute L-8 gate checks and L-9.2
  browser checks in the `browser` package. Set `CHROME_PATH` if Chrome is not on
  `PATH` (Microsoft Edge is used automatically on Windows when Chrome is absent).
- `tc-menu` popovers flip upward when there is not enough space below the
  control; long labels no longer overlap the arrow.

### Included

- `browser` package: chromedp v0.16.0 checks for menu geometry, ellipsis,
  single-row bar, 44 px tap targets, focus return, stacked `Esc`, WCAG AA
  contrast, and stable geometry across themes.
- L-9.3 bite fixtures under `browser/testdata/broken/`.
- `tc.js` `positionPopover()` and `is-flip-up` CSS for menus.

### Not yet included

- Opus API review.

## typecase/v0.0.3 — 2026-09-04

L-8 gates, full catalog coverage, and markup without inline styles.

### For consumers

- Run `go test ./...` from the module root to execute L-8 gate checks in the
  `gate` package. Each gate has a bite test under `gate/testdata/broken/` (L-9.3).
- The live catalog now includes every widget in each declared state. Start it
  with `go run ./cmd/catalog` as before; `/static` still serves the same set
  without `tc.js`.
- Widget markup no longer emits `style=` attributes. Dynamic colours, meter
  width, skeleton width, and crop transforms are carried in `data-tc-*`
  attributes; `tc.js` applies them on load.
- `Input` accepts an optional `Label` field used for `aria-label` when the
  control is not wrapped in `tc-field`.

### Included

- `gate` package: completeness, ownership, node names, typed input,
  enumerations, domain naming, doc format, catalog, no-inline, and a11y checks.
- Catalog entries for all missing widget states; `Shell.Narrow`, `Rail.Layout`.
- `tc.js` `initDynamic()` for data-driven presentation.

### Not yet included

- Opus API review.

## typecase/v0.0.2 — 2026-09-04

Widget families, catalog server, and expanded behavior script.

### For consumers

- Import `carrel.dev/typecase/widget` and call `Render()` on typed widget
  structs (for example `widget.Button{Label: "Save"}.Render()`). Every input
  field is typed; there are no HTML slots.
- Run the live catalog with `go run ./cmd/catalog` from the module root. Open
  `http://localhost:8080/` for the page with `tc.js`, or `/static` for the same
  widgets without behavior scripts.
- Styles now include family sheets (`02-primitives.css` through
  `10-utility.css`) concatenated in `css.Order`. Continue to serve via
  `typecase.CSS()`.
- `js/tc.js` handles menu open/close, overlay `Esc`, scrim dismiss, close
  buttons, and basic focus trapping inside `[data-tc-trap]`. Load after markup
  as before.
- Widget CSS is validated against the public `--tc-*` token surface; do not
  reference undeclared custom properties in overrides.

### Included

- All widget families from L-4: primitives, forms, overlays, panels, lists,
  tables/grids/trees, shell, document, and utility widgets.
- `catalog` package with every widget in each declared state.
- `widget` package tests: render smoke, HTML escaping, catalog completeness,
  per-family CSS token validation.

### Not yet included

- L-9.2 browser checks (chromedp), Opus API review.

## typecase/v0.0.1 — 2026-09-04

First publishable skeleton: module layout, token surface, asset delivery.

### For consumers

- Add `carrel.dev/typecase` as a module dependency (`go 1.27`).
- Call `typecase.CSS()` for the concatenated stylesheet; include it in HTML
  before widget markup.
- Load `typecase.ThemeJS()` in `<head>` **before** any stylesheet or body
  content so the theme is applied before first paint (L-5.3). Load
  `typecase.JS()` at the end of `<body>` for behavior.
- Serve font files from `typecase.Fonts()` at the same relative paths the CSS
  expects (`../fonts/…` from the stylesheet URL).
- Serve `typecase.Sprite()` for the neutral icon sprite; substitute your own
  drawings via the same delivery mechanism when integrating (L-5.1).
- Override appearance only through the documented `--tc-*` custom properties in
  `:root` and `[data-tc-theme="dark"]`. Do not add inline styles or handlers.
  Missing required tokens fail at bundle time, not at runtime with silent
  defaults.

### Included

- MIT-licensed module `carrel.dev/typecase` with `gofmt`, `go test`, and
  `go vet` gates.
- Palette tokens (ink, surface, background, rule, muted, accent, danger) for
  light and dark themes; typography, two densities, radius, and stroke.
- OFL fonts: Literata, Public Sans, Source Code Pro.
- Neutral geometric icon sprite (not application artwork).
- CSS concatenation in declared order (`css.Order`).
- `js/tc.js` (`data-tc-*`) and separate `js/theme.js`.

### Not yet included

- L-8 gate tests beyond token validation, L-9 browser checks (chromedp), and
  full widget catalog acceptance.
