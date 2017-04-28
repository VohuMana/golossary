function onSuccess(data, textStatus, jqXHR ) {
    removeAllChildrenFromNode(document.getElementById("definitionContainer"));
    removeAllChildrenFromNode(document.getElementById("missingDefinitionsContainer"));
    
    var wordsWithNoDefinition = [];
    for (var word in data) {
        if (data[word] !== null) {
            var newDiv = document.createElement('div');
            var header = document.createElement('h3');
            var definition = document.createElement('p');
            header.innerHTML = word;
            definition.innerHTML = data[word];
            newDiv.appendChild(header);
            newDiv.appendChild(definition);
            document.getElementById("definitionContainer").appendChild(newDiv);
        } else {
            wordsWithNoDefinition.push(word);
        }
    }
    
    var noDefinitionContainer = document.getElementById("missingDefinitionsContainer");
    var newHeader = document.createElement('h2');
    newHeader.innerHTML = "Could not find definitions for these words:";
    noDefinitionContainer.appendChild(newHeader);

    var newList = document.createElement('ul');
    for (var index in wordsWithNoDefinition) {
        var listItem = document.createElement('li');
        listItem.innerHTML = wordsWithNoDefinition[index];
        newList.appendChild(listItem);
    }

    noDefinitionContainer.appendChild(newList);
}

function removeAllChildrenFromNode(node) {
    while (node.firstChild) {
        node.removeChild(node.firstChild);
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