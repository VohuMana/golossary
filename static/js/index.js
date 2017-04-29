function onSuccess(data, textStatus, jqXHR ) {
    // Remove all the current definitions that might exist
    removeAllChildrenFromNode(document.getElementById("definitionContainer"));
    removeAllChildrenFromNode(document.getElementById("missingDefinitionsContainer"));
    
    // Alphabetize the words returned
    var alphaWords = alphabetizeWords(data);
    var wordsWithNoDefinition = [];

    // Add definitions for all the words that have definitions.  Save all the words with no definitions.
    for (var index in alphaWords) {
        if (data[alphaWords[index]] !== null) {
            addDefinition(alphaWords[index], data[alphaWords[index]]);
        } else {
            wordsWithNoDefinition.push(alphaWords[index]);
        }
    }
    
    // If there are words that could not be defined add them to the end of the document.
    if (wordsWithNoDefinition.length != 0) {
        addWordsWithNoDefinition(wordsWithNoDefinition);
    }
}

function alphabetizeWords(data) {
    var words = [];

    for (var word in data) {
        words.push(word);
    }

    return words.sort();
}

function addWordsWithNoDefinition(wordsWithNoDefinition) {
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

function addDefinition(word, def) {
    var newDiv = document.createElement('div');
    var header = document.createElement('h3');
    var definition = document.createElement('p');

    header.innerHTML = word;
    definition.innerHTML = def;
    newDiv.appendChild(header);
    newDiv.appendChild(definition);
    document.getElementById("definitionContainer").appendChild(newDiv);
}

function removeAllChildrenFromNode(node) {
    while (node.firstChild) {
        node.removeChild(node.firstChild);
    }
}

function onSubmit() {
    var inputTextArea = document.getElementById("inputTextArea");

    if (inputTextArea.value.length != 0)
    {
        $.ajax({
            type: "POST",
            url: "api/define",
            data: inputTextArea.value,
            success: onSuccess,
            dataType: "json"
        });
    }
}