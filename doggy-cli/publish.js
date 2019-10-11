'use strict';

const publisher = require('@pact-foundation/pact-node');
const path = require('path');
const pjson = require('./package.json');
const opts = {
  pactBroker: process.env.PACT_BROKER_URL,
  tags: [process.env.PACT_TAG],
  pactFilesOrDirs: [path.resolve(process.cwd(), 'pacts')],
  consumerVersion: pjson.version
};

publisher.publishPacts(opts).then(() => done());

