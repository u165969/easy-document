/* prefix index numbers to h3 and h4 texts */
let h = document.querySelectorAll('h3, h4');
let i = 0;
let j = 0;
for(let k = 0; k < h.length; k++) {
  let value = h[k]
  if(value.tagName == 'H3') {
    i++;
    j = 0;
    value.textContent = i + ' ' + value.textContent;
  } else if(value.tagName == 'H4') {
    j++;
    value.textContent = i + '.' + j + ' ' + value.textContent;
  }
}

/* prefix h1 text to title */
let intro = document.getElementsByTagName('h1');
let title = document.getElementsByTagName('title');
title[0].textContent = intro[0].textContent + ' - ' + title[0].textContent;
