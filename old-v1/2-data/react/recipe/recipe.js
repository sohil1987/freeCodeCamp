var box = (function () {
  'use strict';
  /* code here */

  var basicRecipes = [{
    title: 'Cookies',
    ingredients: ['Cup shortening', 'Peanut Butter', 'Milk', 'Eggs',
      'Vanilla', 'Sugar'
    ],
    image: 'http://cdn.pcwallart.com/images/chocolate-chip-cookies-and-milk-wallpaper-2.jpg'
  }, {
    title: 'Spaghetti',
    ingredients: ['Noodles', 'Tomato Sauce', 'Meatballs', 'Onion'],
    image: 'http://cdn-image.myrecipes.com/sites/default/files/styles/300x300/public/image/recipes/ay/06/spaghetti-meatballs-ay-1875344-x.jpg?itok=rou85Q3g'
  }, {
    title: 'Rice Pudding',
    ingredients: ['White Rice', 'Milk', 'Sugar', 'Salt'],
    image: 'http://cdn-image.myrecipes.com/sites/default/files/styles/300x300/public/image/recipes/rs/2007/rice-pudding-rs-1582788-x.jpg?itok=vFbNvIzh'
  }];

  function init () {
    console.log('Inicio');
    populateTable(getData());
    addEvents();
  }

  function addEvents () {
    var actions = document.getElementsByClassName('btn');
    for (var i = 0; i < actions.length; i++) {
      actions.item(i).addEventListener('click', click);
    }
  }

  function click (ev) {
    switch (ev.target.id) {
      case 'addRecipe':
        addRecipe();
        break;
      case 'delRecipe':
        delRecipe(ev);
        break;
      case 'editRecipe':
        editRecipe(ev);
        break;
      default:
        console.log(ev.target, ' -- Not recognized event');
    }
  }

  function addRecipe () {
    createForm(-1);
  }

  function delRecipe (ev) {
    var recipes = JSON.parse(localStorage.recipes);
    recipes.splice(ev.target.value, 1);
    localStorage.recipes = JSON.stringify(recipes);
    init();
  }

  function editRecipe (ev) {
    createForm(ev.target.value);
  }

  function createForm (value) {
    var recipes = JSON.parse(localStorage.recipes);
    var res = '';
    var aux = {};
    if (recipes[value]) {
      aux.title = recipes[value].title;
      aux.image = recipes[value].image;
      aux.ingredients = recipes[value].ingredients;
      res = getRecipeForm(aux);
    } else {
      res = getRecipeForm();
    }
    document.getElementById('recipeForm').innerHTML = res;
    var modal = document.getElementById('recipeForm');
    modal.style.display = 'block';
    window.addEventListener('click', function (ev) {
      if (ev.target === modal) modal.style.display = 'none';
    });
    var span = document.getElementsByClassName('close')[0];
    span.addEventListener('click', function () {
      modal.style.display = 'none';
    });
    var saveAndClose = document.getElementById('saveAndClose');
    saveAndClose.addEventListener('click', function (ev) {
      var c1 = document.getElementById('recipeName').value;
      var c2 = document.getElementById('recipeImageUrl').value;
      var c3 = document.getElementById('recipeIngredients').value;
      urlIsOk(c2, recipes, value);
      if (c2 === '') {
        document.getElementById('recipeImageUrl').value =
          './images/photoNot.png';
      }
      if (c1 !== '' && c3 !== '') {
        modal.style.display = 'none';
        if (recipes[value]) {
          recipes[value].title = document.getElementById(
            'recipeName').value;
          recipes[value].image = document.getElementById(
            'recipeImageUrl')
            .value;
          recipes[value].ingredients = document.getElementById(
            'recipeIngredients').value;
          localStorage.recipes = JSON.stringify(recipes);
          init();
        } else if (value === -1) {
          recipes.push({});
          recipes[recipes.length - 1].title = document.getElementById(
            'recipeName').value;
          recipes[recipes.length - 1].image = document.getElementById(
            'recipeImageUrl').value;
          recipes[recipes.length - 1].ingredients = document.getElementById(
            'recipeIngredients').value;
          localStorage.recipes = JSON.stringify(recipes);

          init();
        }
      }
    });
  }

  function urlIsOk (imageUrl, recipes, value) {
    var imageData = new Image();
    /*imageData.onload = function () {
      return true
    };*/
    imageData.onerror = function () {
      if (recipes[value]) {
        recipes[value].image = './images/photoNot.png';
        localStorage.recipes = JSON.stringify(recipes);
        init();
      }
      if (value === -1) {
        recipes[recipes.length - 1].image = './images/photoNot.png';
        localStorage.recipes = JSON.stringify(recipes);
        init();
      }
    };
    imageData.src = imageUrl;
  }

  function getRecipeForm (recipe) {
    if (recipe === undefined) {
      recipe = {};
      recipe.action = 'Create New Recipe';
      recipe.title = '';
      recipe.image = '';
      recipe.ingredients = '';
      recipe.titleP = 'Enter recipe name';
      recipe.imageP = 'Enter URL image';
      recipe.ingredientsP = 'Enter ingredients separated by commas';
    } else {
      recipe.action = 'Edit Recipe';
      recipe.titleP = 'Enter recipe name';
      recipe.imageP = 'Enter URL image';
      recipe.ingredientsP = 'Enter ingredients separated by commas';
    }
    var res = '';
    res += '<div class="modal-content">';
    res += '<div class="modal-header">';
    res += '<span class="close">×</span>';
    res += '<h4 id="actionType">' + recipe.action + '</h4>';
    res += '</div>';
    res += '<div class="modal-body">';
    res += '<form>';
    res += '<div class="form-group">';
    res += '<label for="recipeName">Recipe</label>';
    res +=
      '<input type="text" class="form-control" id="recipeName" placeholder="' +
      recipe.titleP + '" value="' + recipe.title + '">';
    res += '</div>';

    res += '<div class="form-group">';
    res += '<label for="recipeImageUrl">URL image</label>';
    res +=
      '<input type="text" class="form-control" id="recipeImageUrl" placeholder="' +
      recipe.imageP + '" value="' + recipe.image + '">';
    res += '</div>';

    res += '<div class="form-group">';
    res += '<label for="recipeIngredients">Ingredients</label>';
    res +=
      '<input type="text" class="form-control" id="recipeIngredients" placeholder="' +
      recipe.ingredientsP + '" value="' + recipe.ingredients + '">';
    res += '</div>';
    res += '</form>';
    res += '</div>';
    res +=
      '<div class="modal-footer"><h4 id="saveAndClose">Save and Close</h4></div>';
    res += '</div>';
    return res;
  }

  function populateTable (recipes) {
    // console.log('Recipes in memory --> ', recipes.length)
    var res = '';
    for (var i = 0; i < recipes.length; i++) {
      res +=
        '<div class="card text-xs-center col-xs-6 col-sm-4 col-md-3">';
      res += '<div class="recipe">';
      res += '<img class="recipeImg" alt="" src="' +
        recipes[i].image +
        '" ><br>';
      res +=
        '<p class="recipeText"><span class="maxLine"><strong>' + recipes[i]
          .title +
        '</strong></span></p></div>';
      res += '<div class="btnGroup btn-group-sm" role="group">';
      res +=
        '<button id="delRecipe" value="' + i +
        '" type="button" class="btn btn-warning">Del</button>';
      res +=
        '<button id="editRecipe" value="' + i +
        '" type="button" class="btn btn-info">Edit</button>';
      res += '</div></div>';
    }
    // onerror="this.onerror=null;this.src="./images/photoNot.png"
    document.getElementById('dataTable').innerHTML = res;
  }

  function getData () {
    // console.log(localStorage.recipes)
    if (localStorage.recipes === undefined) {
      localStorage.recipes = JSON.stringify(basicRecipes);
    // localStorage.setItem('recipes', JSON.stringify(basicRecipes))
    }

    // para pruebas que no se quede vacia de recetas la cosa
    if (JSON.parse(localStorage.recipes).length === 0) {
      localStorage.recipes = JSON.stringify(basicRecipes);
    }
    return JSON.parse(localStorage.recipes);
  // return JSON.parse(localStorage.getItem('recipes'))
  // clear localStorage // localStorage.clear()
  }

  window.addEventListener('load', init);
}());
