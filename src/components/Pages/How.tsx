import React, { Component } from 'react';
import { Link } from 'react-router-dom';

export class How extends Component {
    public constructor(props: any) {
        super(props);
    }

    public render() {
        return(
            <div className="page-wrap examples">
                <h2>How it Works</h2>

                <p>
                    Divi's Atomic Swaps are based off of
                    <a href="https://github.com/decred/atomicswap" target="decred" className="text-link">
                        <i className="fab fa-github"></i>
                        the original implementation by Decred.
                    </a>
                    And, the base
                    <Link className="text-link" to="/api">
                        <i className="fas fa-book-open"></i>
                        API Reference
                    </Link>
                    has the ability to use similar granular functionality to the original API
                    in the extended portion of the reference.
                </p>

                <p>
                    Divi's core intention was to create a production use case for Atomic Swaps.
                    We have created an API where 1 API call via command line or HTTP POST can broker
                    an Atomic Swap between Bitcoin and Divi.
                </p>

                <h3>The original implementation</h3>

                <p>
                    In the original implementation by Decred.
                    An Atomic Swap was executed procedurally by the user.
                    This same functionality is also part of our extended API if you need more granular control
                    over the process.
                </p>

                <ol>
                   <li>User 1 initiates an Atomic Swap on chain 1</li>
                   <li>User 1 gives the secret hash to User 2</li>
                   <li>User 2 create a participating atomic swap on chain 2</li>
                   <li>User 1 and 2 redeem the atomic swap contracts on their respective chains</li>
                </ol>

                <h3>Divi's implementation</h3>

                <p>
                    We have improved this implementation by automating the entire process for atomic swaps.
                    Simply just send a HTTP request to initiate a swap to execute an Atomic Swap.
                    Just be mindful that blocks need to be mined for the atomic swap to execute.
                </p>

                <ol>
                    <li>User sends a post request to initiate an atomic swap</li>
                    <li>The Atomic Swap node automates the entire process.</li>
                </ol>

                <h3>Getting Started</h3>

                <p>
                    Now that you know how to get started with Atomic Swaps. You can go to our
                    <Link className="text-link" to="/examples">
                        <i className="fas fa-graduation-cap"></i>
                        Examples
                    </Link>
                    section which will teach you how to setup a node and execute swaps.
                    We hope you enjoy this free open source implementation.
                </p>
            </div>
        )
    }
}
