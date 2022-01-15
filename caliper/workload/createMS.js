'use strict';

const { WorkloadModuleBase } = require('@hyperledger/caliper-core');

class MyWorkload extends WorkloadModuleBase {
    constructor() {
        super();
        this.txIndex = -1;
        this.medName = ['aspirin', 'vicodin', 'synthroid', 'delasone', 'amoxil', 'neurontin', 'zestril', 'lipitor', 'glucophage', 'zofran', 'ibuprofen']
        this.disease = ['Pain management', 'Thyroid deficiency', 'Arthritis', 'Bacterial infections', 'Seizures', 'Blood pressure', 'High cholesterol', 'Type 2 diabetes', 'Fever']
    }

    /**
    * Initialize the workload module with the given parameters.
    * @param {number} workerIndex The 0-based index of the worker instantiating the workload module.
    * @param {number} totalWorkers The total number of workers participating in the round.
    * @param {number} roundIndex The 0-based index of the currently executing round.
    * @param {Object} roundArguments The user-provided arguments for the round from the benchmark configuration file.
    * @param {ConnectorBase} sutAdapter The adapter of the underlying SUT.
    * @param {Object} sutContext The custom context object provided by the SUT adapter.
    * @async
    */
    async initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext) {
        await super.initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext);
    }

    async randomDate() {
        let year = (2020 + Math.floor(Math.random() * 10)).toString()
        let month = Math.floor((Math.random() * 12) + 1)
        let day = Math.floor((Math.random() * 31) + 1)
        if (month < 10) month = '0' + month;
        if (day < 10) day = '0' + day;
        return [year, month, day].join('.');
    }

    async submitTransaction() {
        this.txIndex++;

        let medName = this.medName[this.txIndex % this.medName.length];
        const medNumber = `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
        let disease = this.disease[this.txIndex % this.disease.length];
        let date = this.randomDate()

        let price = Math.floor(Math.random() * 1000).toString() // random number between 0 and 1000

        const request = {
            contractId: this.roundArguments.contractId,
            contractFunction: 'Issue',
            invokerIdentity: 'alice',
            contractArguments: [medName, medNumber, disease, date, price],
            readOnly: false
        };

        await this.sutAdapter.sendRequests(request);
    }

}

function createWorkloadModule() {
    return new MyWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;