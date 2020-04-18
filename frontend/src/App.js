import React from 'react';

import { HashRouter as Router, Switch, Route } from 'react-router-dom';
import Compare from './Components/Compare'
import Ranking from './Components/Ranking'

import { Main } from '@aragon/ui'


import "bootstrap/dist/css/bootstrap.min.css";
import "shards-ui/dist/css/shards.min.css"

function App() {
  return (
    <Main>
      <Router>
        <Switch>
          <Route path='/ranking/' children={<Ranking />} />
          <Route path='/' children={<Compare />} />
        </Switch>
      </Router>
    </Main>
  );
}

export default App;
