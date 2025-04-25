import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Sidebar from './components/Sidebar';
import ChatWindow from './components/ChatWindow';

function App() {
  return (
    <Router>
      <Sidebar />
      <Switch>
        <Route path="/room1" component={ChatWindow} />
        <Route path="/room2" component={ChatWindow} />
      </Switch>
    </Router>
  );
}

export default App;
