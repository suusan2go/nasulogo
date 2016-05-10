import React, { Component } from 'react';
import Helmet from 'react-helmet';

export default class App extends Component {
  render() {
    return(
      <div>
        <header>
          <div className="navbar navbar-default" role="navigation">
            <div className="container">
              <div className="navbar-header">
                <button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#gloabal-header" aria-expanded="false">
                  <span className="sr-only">Toggle navigation</span>
                  <span className="icon-bar"></span>
                  <span className="icon-bar"></span>
                  <span className="icon-bar"></span>
                </button>
                <a className="navbar-brand" href="/home">Nasulog</a>
              </div>
              <div className="collapse navbar-collapse" id="gloabal-header">
                <p className="navbar-text navbar-right">
                  <a href="/auth//login/google">ログイン/新規登録</a>
                </p>
              </div>
            </div>
          </div>
        </header>
        <Helmet title='Go + React + Redux = rocks!' />
        {this.props.children}
      </div>
    )
  }

}
