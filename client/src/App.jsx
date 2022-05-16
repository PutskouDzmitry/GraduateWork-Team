import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import Header from "./components/Header";
import Main from "./components/Main";
import ACSparser from "./components/ACSparser";
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
            <Header withLoginButton={false} />
            <LogInForm type={"login"} />
            <Footer />
          </Route>
          <Route path="/register">
            <Header withLoginButton={false} />
            <LogInForm type={"signin"} />
            <Footer />
          </Route>
          <Route path="/acs-parser">
            <Header withLoginButton={true} />
            <Sidebar />
            <ACSparser />
            <RouterSettings />
            <Footer />
          </Route>
          <Route path="/">
            <Header withLoginButton={true} />
            <Sidebar />
            <Main />
            <RouterSettings />
            <Footer />
          </Route>
        </Switch>
      </>
    </Router>
  );
}

export default App;
