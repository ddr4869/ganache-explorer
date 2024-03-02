const tomtoken = artifacts.require("TomToken");

module.exports = function(deployer) {
    deployer.deploy(tomtoken);
}
