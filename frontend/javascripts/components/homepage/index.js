import React, { Component } from 'react';
import Helmet from 'react-helmet';

export default class Homepage extends Component {
  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    // Load here any data.
    callback(); // this call is important, don't forget it
  }
  /*eslint-enable */

  render() {
    return (
      <div className="container">
        <Helmet
          title="Home page"
          meta={[{
            property: 'og:title',
            content: 'Golang Isomorphic React/Hot Reloadable/Redux/Css-Modules Starter Kit' }]}
        />
        <div className="page-header">
          <h1>みんなのポエム一覧</h1>
        </div>
        <p>
          <a className="btn btn-primary btn-lg" href="/poems/new">ポエムを作成する</a>
        </p>
        <ul className="list-group">
          <div className="row">
            <li className="list-group-item col-md-4">
              <a href="/poems/2992">
                <div className="medila">
                  <div className="media-left">
                    <img className="img-circle" width="60" src="" alt="person" />
                  </div>
                  <div className="media-body">
                    <div className="media-heading">
                      <span className="link-text">Warning</span>
                    </div>
                    <div className="media-content">
                      sample
                    </div>
                    <div className="text-right">
                      2016/04/28 18:47
                    </div>
                  </div>
                </div>
              </a>
            </li>
          </div>
        </ul>
      </div>
    );
  }
}
