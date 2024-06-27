export function contentEditableFocus(node: HTMLElement) {
  const selectedText = window.getSelection();
  if (!selectedText || !node.contentEditable) return;
  const selectedRange = document.createRange();
  selectedRange.setStart(node.childNodes[0], node.textContent?.length ?? 0);
  selectedRange.collapse(true);
  selectedText.removeAllRanges();
  selectedText.addRange(selectedRange);
  node.focus();
}

export function contentEditableMacroFocus(node: HTMLElement) {
  setTimeout(() => contentEditableFocus(node));
}
