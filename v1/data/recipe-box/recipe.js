/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  const notImage = '/freecodecamp/assets/images/photoNot.png';

  const basicRecipes = [
    {
      title: 'Cookies',
      ingredients: ['Cup shortening', 'Peanut Butter', 'Milk', 'Eggs',
        'Vanilla', 'Sugar'],
      image: 'http://cdn.pcwallart.com/images/chocolate-chip-cookies-and-milk-wallpaper-2.jpg'
    }, {
      title: 'Blueberry',
      ingredients: ['Vodka', 'Blueberry'],
      image: 'http://img.huffingtonpost.com/asset/scalefit_600_noupscale/55c96a571d00002f001446b7.jpeg'
    }, {
      title: 'Spaghetti',
      ingredients: ['Noodles', 'Tomato Sauce', 'Meatballs', 'Onion'],
      image: 'http://cdn-image.myrecipes.com/sites/default/files/styles/300x300/public/image/recipes/ay/06/spaghetti-meatballs-ay-1875344-x.jpg?itok=rou85Q3g'
    }, {
      title: 'Food with no Photo but with a very very very very long name.',
      ingredients: ['Potatoes'],
      image: notImage
    }, {
      title: 'Rice Pudding',
      ingredients: ['White Rice', 'Milk', 'Sugar', 'Salt'],
      image: 'http://cdn-image.myrecipes.com/sites/default/files/styles/300x300/public/image/recipes/rs/2007/rice-pudding-rs-1582788-x.jpg?itok=vFbNvIzh'
    }, {
      title: 'Vegetarian pizza',
      ingredients: ['Pizza', 'Vegetables'],
      image: 'https://s-media-cache-ak0.pinimg.com/736x/b4/b0/7b/b4b07b7e8d77fa0990a33385b8675365--tone-it-up-vegetarian-recipes-flavorful-vegetarian-recipes.jpg'
    }, {
      title: 'Pesto Chicken',
      ingredients: ['Chicken', 'Vegetables'],
      image: 'https://s-media-cache-ak0.pinimg.com/736x/8d/11/00/8d11005c1f137b8f58e1f0db6871fb58.jpg'
    }, {
      title: ' Secret Cookies of Dart Vader',
      ingredients: ['Cup shortening', 'Peanut Butter', 'Milk', 'Eggs',
        'Vanilla', 'Sugar'],
      image: 'http://4.bp.blogspot.com/-_qRAJrtalU0/TquMymJPBUI/AAAAAAAAAk8/1AotB3Oy4wM/s320/IMG_5561.jpg'
    }];

  function init () {
    console.log('Init Recipe Box');
    populateTable(getData());
    addEvents();
  }

  function addEvents () {
    const actions = document.getElementsByClassName('btn');
    for (let i = 0; i < actions.length; i++) {
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
    let recipes = JSON.parse(localStorage.recipes);
    recipes.splice(ev.target.value, 1);
    localStorage.recipes = JSON.stringify(recipes);
    init();
  }

  function editRecipe (ev) {
    createForm(ev.target.value);
  }

  function createForm (value) {
    let recipes = JSON.parse(localStorage.recipes);
    let res = '';
    let aux = {};
    if (recipes[value]) {
      aux.title = recipes[value].title;
      aux.image = recipes[value].image;
      aux.ingredients = recipes[value].ingredients;
      res = getRecipeForm(aux);
    } else {
      res = getRecipeForm();
    }
    document.getElementById('recipeForm').innerHTML = res;
    let modal = document.getElementById('recipeForm');
    modal.style.display = 'block';
    window.addEventListener('click', function (ev) {
      if (ev.target === modal) modal.style.display = 'none';
    });
    let span = document.getElementsByClassName('close')[0];
    span.addEventListener('click', function () {
      modal.style.display = 'none';
    });
    let saveAndClose = document.getElementById('saveAndClose');
    saveAndClose.addEventListener('click', function (ev) {
      let c1 = document.getElementById('recipeName').value;
      let c2 = document.getElementById('recipeImageUrl').value;
      let c3 = document.getElementById('recipeIngredients').value;
      urlIsOk(c2, recipes, value);
      if (c2 === '') {
        document.getElementById('recipeImageUrl').value = notImage;
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
    const imageData = new Image();
    /*imageData.onload = function () {
      return true
    };*/
    imageData.onerror = function () {
      if (recipes[value]) {
        recipes[value].image = notImage;
        localStorage.recipes = JSON.stringify(recipes);
        init();
      }
      if (value === -1) {
        recipes[recipes.length - 1].image = notImage;
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
    let res = `
    <div class="modal-content">
    
    <div class="modal-header">
      <span class="close">×</span>
      <h4 id="actionType"> ${recipe.action}</h4>
    </div>
    
    <div class="modal-body">
      <form>
      <div class="entrie">
        <label for="recipeName">Recipe</label>
          <input type="text" class="for-control" id="recipeName" placeholder="${recipe.titleP}" value="${recipe.title}">
      </div>
      <div class="entrie">
        <label for="recipeImageUrl">URL image</label>
        <input type="text" class="for-control" id="recipeImageUrl" placeholder="${recipe.imageP}" value="${recipe.image}">
      </div>
      <div class="entrie">
        <label for="recipeIngredients">Ingredients</label>
        <input type="text" class="for-control" id="recipeIngredients" placeholder="${recipe.ingredientsP}" value="${recipe.ingredients}">
      </div>
      </form>
    </div>

    <div class="modal-footer">
      <h4 id="saveAndClose">Save and Close</h4>
    </div>
    
    </div>
    `;
    return res;
  }

  function populateTable (recipes) {
    // console.log('Recipes in memory --> ', recipes.length)
    let res = '';
    for (let i = 0; i < recipes.length; i++) {
      res +=
        `<div class="card">
          <div class="recipe">
            <img class="recipeImg" alt="" src="${recipes[i].image}">
            <div class="recipeText">${recipes[i].title}</div>
          </div>
          <div class="controls">
            <div class="action">
              <button id="delRecipe" class="btn" value=${i}>Del</button>
            </div>
            <div class="action">
              <button id="editRecipe" class="btn" value=${i}>Edit</button>
            </div>
          </div>
        </div>`;
    }
    // console.log(res)
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

  return {
    init: init
  };
}());

window.addEventListener('load', app.init);
