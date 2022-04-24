import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import Header from "./components/Header";
import Main from "./components/Main";
import Footer from "./components/Footer";
import Sidebar from "./components/Sidebar";
import LogInForm from "./components/LogInForm";
import RouterSettings from "./components/RouterSettings";

import "./styles/common.scss";

function App() {
  return (
    <Router>
      <>
        <Switch>
          <Route path="/login">
            <LogInForm />
          </Route>
          <Route path="/">
            <Sidebar />
            <Header />
            <Main />
            <Footer />
            <RouterSettings />
          </Route>
        </Switch>
      </>
    </Router>
  );
}

export default App;
