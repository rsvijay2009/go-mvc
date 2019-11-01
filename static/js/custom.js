function deleteNote(noteId) {
    let choice = confirm("Are you sure to delete this note ?");

    if (choice == true) {
        location.href = "notepad/delete/" + noteId;
        return false
    }
}