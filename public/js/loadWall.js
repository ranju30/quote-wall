function emptyField() {
    $("#name").val("");
    $("#textArea").val("");
}
function store() {
    var toSaveName = $("#name").val();
    var toSaveQuote = $("#textArea").val();
    if (toSaveName != "" && toSaveQuote != "") {
        emptyField();
        $.post("/store", {name: toSaveName, quote: toSaveQuote}, function (res, err) {
            $('#updateWall').append('<b>' + toSaveName + '</b>' + ' : ' + toSaveQuote + '<br>');
        })
    }
}

function allQuotes() {
    $.get("/store", function (res, err) {
        $('#output').html(res);
    })
}
// window.load = allQuotes();


