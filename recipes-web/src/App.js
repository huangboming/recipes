import React from 'react';
import './App.css';
import Recipe from './Recipe';


class App extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      recipes: []
    }

    this.getRecipes = this.getRecipes.bind(this);
    this.getRecipes();
  }

  render() {
    return (<div>
      {this.state.recipes.map((recipe, index) => (
        <Recipe recipe={recipe} />
      ))}
    </div>);
  }

  async getRecipes() {
    return fetch('http://localhost:8080/recipes', {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
      },
    })
      .then(response => response.json())
      .then(data => this.setState({ recipes: data }));
  }
}

export default App;