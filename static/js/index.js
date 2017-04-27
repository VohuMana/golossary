function onSuccess(data, textStatus, jqXHR ) {
    for (var word in data) {
        if (data[word] !== null) {
            var newDiv = document.createElement('div');
            var header = document.createElement('h3');
            var definition = document.createElement('p');
            header.innerHTML = word;
            definition.innerHTML = data[word];
            newDiv.appendChild(header);
            newDiv.appendChild(definition);
            document.getElementById("mainContainer").appendChild(newDiv);
        }
    }
}

function onSubmit() {
    var inputTextArea = document.getElementById("inputTextArea");

    $.ajax({
        type: "POST",
        url: "api/define",
        data: inputTextArea.value,
        success: onSuccess,
        dataType: "json"
    });
}