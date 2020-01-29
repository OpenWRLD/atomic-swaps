import React, { Component } from 'react';
import { UnControlled as CodeMirror } from 'react-codemirror2';

export const SwapHTTP = `$ curl -d \\
'{ \\
    "baseAmount": "1.0000", \\
    "baseAddress": "1A1zsatoshiefi2DMPpoopamotomv7DivfNa", \\
    "swapAmount": "1.0000", \\
    "swapAddress": "1A1zsatoshiefi2DMPpoopamotomv7DivfNa", \\
}' \\
$ -H "Content-Type: application/json" \\
$ -X POST http://localhost:3000/v1/swap`;

export const SwapJSON = `{
    "status": "pending",
    "swapId": "5e27b27a6f97e18a45726af2"
}`;

export class SwapExample extends Component {
    public constructor(props: any) {
        super(props);
        this.state = { menu: false };
    }

    public render() {
        return(
            <div className="example-wrap" id="e10">
                <h2>Executing an Atomic Swap</h2>
                <div className="example">
                    <div className="example-content">
                        <div className="example-text">
                            <p>
                                With one post request, you will be able to execute an atomic swap.
                            </p>
                            <p>
                                The base address is the participant's address. This is the Base RPC sending money to the
                                participant.
                            </p>
                            <p>
                                The swap address is the initiator's address. This is the Swap RPC sending money to the
                                initiator.
                            </p>
                        </div>
                        <CodeMirror
                            value={SwapHTTP}
                            options={{
                                mode: 'shell',
                                theme: 'material',
                                lineNumbers: true
                            }}
                        />
                        <div className="example-text">
                            <p>
                                Once the process starts. You will receive a response with your swap id. This swap id can be used
                                to query the atomic swap server to check the status of your swap.
                            </p>
                        </div>
                        <CodeMirror
                            value={SwapJSON}
                            options={{
                                mode: 'shell',
                                theme: 'material',
                                lineNumbers: true
                            }}
                        />
                    </div>
                </div>
            </div>
        );
    }
}
