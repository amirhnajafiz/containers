function editor(element, input) {
    element.innerHTML = markdown.toHTML(input);

    return element
}

export default editor;