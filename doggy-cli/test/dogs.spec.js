'use strict'

const expect = require('chai').expect;
const path = require('path');
const { Pact, Matchers } = require('@pact-foundation/pact');
const { getDog, saveDog } = require('../index');
const { integer, string } = Matchers;

describe('The Dog API', () => {
  let url = 'http://localhost';
  const port = 8992;

  const provider = new Pact({
    port: port,
    log: path.resolve(process.cwd(), 'logs', 'mockserver-integration.log'),
    dir: path.resolve(process.cwd(), 'pacts'),
    spec: 2,
    consumer: 'doggy-cli',
    provider: 'petstore',
    pactfileWriteMode: 'merge',
  });

  const dogExample = {
    id: 88,
    name: 'Chico',
    type: 'Shiba Inu'
  };
  const EXPECTED_BODY = {
    id: integer(dogExample.id),
    name: string(dogExample.name),
    type: string(dogExample.type)
  };

  // Setup the provider
  before(() => provider.setup());

  // Write Pact when all tests done
  after(() => provider.finalize());

  // verify with Pact, and reset expectations
  afterEach(() => provider.verify());

  describe('get /dog/88', () => {
    before(done => {
      const interaction = {
        state: 'there is a dog with an id 88',
        uponReceiving: 'a request to get dog id 88',
        withRequest: {
          method: 'GET',
          path: '/dogs/88',
          headers: {
            Accept: 'application/json;charset=UTF-8',
          },
        },
        willRespondWith: {
          status: 200,
          headers: {
            'Content-Type': 'application/json;charset=UTF-8',
          },
          body: EXPECTED_BODY,
        },
      }
      provider.addInteraction(interaction).then(() => {
        done()
      })
    });

    it('returns the correct response', done => {
      const urlAndPort = {
        url: url,
        port: port,
      };
      getDog(urlAndPort, 88).then(response => {
        expect(response.data).to.deep.eq(dogExample)
        done()
      }, done);
    });
  });

  describe('post /dogs', () => {
    const dog = {
      name: 'Chico',
      type: 'Shiba Inu'
    };
    before(done => {
      const interaction = {
        state: 'creating a Shiba Inu dog whose name is Chico',
        uponReceiving: 'a request to add a Shiba Inu dog whose name is Chico',
        withRequest: {
          method: 'POST',
          path: '/dogs',
          headers: {
            'Content-Type': 'application/json;charset=UTF-8',
            Accept: 'application/json;charset=UTF-8',
          },
          body: dog
        },
        willRespondWith: {
          status: 201,
          headers: {
            'Content-Type': 'application/json;charset=UTF-8',
          },
          body: EXPECTED_BODY,
        },
      };
      provider.addInteraction(interaction).then(() => {
        done()
      });
    });

    it('returns the correct response', done => {
      const urlAndPort = {
        url: url,
        port: port,
      };
      saveDog(urlAndPort, dog).then(response => {
        expect(response.data).to.deep.eql(dogExample)
        done()
      }, done);
    });
  });
});
