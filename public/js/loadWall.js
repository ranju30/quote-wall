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
            $('#updateWall').append('<b>' + toSaveName + '</b>' + ' : ' + '<i>' + toSaveQuote + '</i>' + '<br>');
        })
    }
}

function allQuotes() {
    $.get("/store", function (res, err) {
        $('#updateWall').html(res);
    })
}
window.load = allQuotes();


