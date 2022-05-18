import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import Header from "./components/Header";
import Main from "./components/Main";
import ACSparser from "./components/ACSparser";
import MobileParser from "./components/MobileParser";
import AcrylicParser from "./components/AcrylicParser";
import Footer from "./components/Footer";
import Sidebar from "./components/Sidebar";
import LogInForm from "./components/LogInForm";
import RouterSettings from "./components/RouterSettings";
import LoaderModal from "./components/LoaderModal";

import "./styles/common.scss";

function App() {
  return (
    <Router>
      <>
        <Switch>
          <Route path="/login">
            <Header withLoginButton={false} />
            <LogInForm type={"login"} />
            <LoaderModal />
            <Footer />
          </Route>
          <Route path="/register">
            <Header withLoginButton={false} />
            <LogInForm type={"signin"} />
            <LoaderModal />
            <Footer />
          </Route>
          <Route path="/acs-parser">
            <Header withLoginButton={true} />
            <Sidebar />
            <ACSparser />
            <RouterSettings />
            <LoaderModal />
            <Footer />
          </Route>
          <Route path="/mobile-parser">
            <Header withLoginButton={true} />
            <Sidebar />
            <MobileParser />
            <RouterSettings />
            <LoaderModal />
            <Footer />
          </Route>
          <Route path="/acrylic-parser">
            <Header withLoginButton={true} />
            <Sidebar />
            <AcrylicParser />
            <RouterSettings />
            <LoaderModal />
            <Footer />
          </Route>
          <Route path="/">
            <Header withLoginButton={true} />
            <Sidebar />
            <Main />
            <RouterSettings />
            <LoaderModal />
            <Footer />
          </Route>
        </Switch>
      </>
    </Router>
  );
}

export default App;
