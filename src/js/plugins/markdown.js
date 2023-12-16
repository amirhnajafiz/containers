function editor(element, input) {
    element.innerHTML = marked.parse(input);

    return element
}

export default editor;
