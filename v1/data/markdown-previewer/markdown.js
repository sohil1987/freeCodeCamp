const app = (function () {
  'use strict';
  /* code here */

  function init () {
    console.log('Init Markdown Previewer');
    window.addEventListener('resize', adjustSize);
    document.getElementById('editor').addEventListener('input', previewText);
    getPlaceholder();
    adjustSize();
  }

  function getPlaceholder () {
    document.getElementById('editor').placeholder = text;
    document.getElementById('preview').innerHTML = marked(text);
  }

  function previewText (ev) {
    document.getElementById('preview').innerHTML = marked(ev.target.value);
  }

  function adjustSize () {
    let h = document.getElementById('preview').style.height | window.innerHeight;
    document.getElementById('editor').style.height = h - 60 + 'px';
  }

  return {
    init: init
  };
}());

let text =
`Its very easy to make some words **bold** and other words *italic* with Markdown  

You can even [link to my Freecodecamp projects!](https://brusbilis.com/freecodecamp)

Sometimes you want numbered lists:  

1. One  
2. Two  
3. Three  
Sometimes you want bullet points:

* Start a line with a star
* Profit!

Alternatively,

- Dashes work just as well
- And if you have sub points, put two spaces before the dash or star:
  - Like this
  - And this

If you want to embed images, this is how you do it:
Markdown Logo  
![Markdown Logo](/assets/icons/markdown.png)

# Structured documents

Sometimes its useful to have different levels of headings to structure your documents. Start lines with a # to create headings. Multiple ## in a row denote smaller heading sizes.

### This is a third-tier heading

You can use one # all the way up to ###### six for different heading sizes.

If youd like to quote someone, use the > character before the line:

> Coffee. The finest organic suspension ever devised... I beat the Borg with it.
> - Captain Janeway

GitHub supports many extras in Markdown that help you reference and link to people. If you ever want to direct a comment at someone, you can prefix their name with an @ symbol: Hey @kneath — love your sweater!

But I have to admit, tasks lists are my favorite:

- [x] This is a complete item
- [ ] This is an incomplete item

When you include a task list in the first comment of an Issue, you will see a helpful progress bar in your list of issues. It works in Pull Requests, too!

[Doc from Github](https://guides.github.com/features/mastering-markdown/#examples)
`;

text = text + text + text + text + text;

addEventListener('load', app.init());
