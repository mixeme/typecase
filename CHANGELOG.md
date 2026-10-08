# Changelog

All notable changes to Typecase are documented here. Versioning follows
[semantic versioning](https://semver.org/). This file tells consumers what to
do, not what the authors did.

## Unreleased

### Breaking

- `Badge.Icon` is gone: no frame draws a badge with an icon since the
  «linked» badge left the tags column of the notes list (Carrel's task 105).
- `TickRow.Trail` no longer takes a `Button`: a refused source has no Retry
  in the rail any more (the owner's decision of 02.10, below).
- `Sec` no longer takes a `Table`: no frame puts one in a section since the
  attendees became a `SecList` with `Columns` (Carrel's task 77).

### Added

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
