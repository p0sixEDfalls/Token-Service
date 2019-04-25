//This is simply ERC20 token example

pragma solidity ^0.5.0;

contract ERC20Token {

	uint8 constant public decimals = 0;
    uint256 public totalTokenSupply = 10**6; // 1 million tokens, 0 decimal places
    string constant public name = "Simple ERC20 Token";
    string constant public symbol = "SET";
    address owner;

    mapping (address => uint256) private balances;

    modifier onlyOwner() 
  	{
    	require(msg.sender == owner, "access denied");
    	_;
	}

    constructor() public {
    	owner = msg.sender;
    	balances[owner] = totalTokenSupply;
    }

    function totalSupply() view public returns (uint256 supply) {
    	return totalTokenSupply;
    }

    function balanceOf(address _owner) view public returns (uint256 balance) {
    	return balances[_owner];
    }

    function transfer(address _to, uint256 _value) public onlyOwner returns (bool success) {
    	 if (balances[msg.sender] >= _value && balances[_to] + _value >= balances[_to]) {
            balances[msg.sender] -= _value;
            balances[_to] += _value;
            return true;
        } else { return false; }
    }

    function transferFrom(address _from, address _to, uint256 _value) public onlyOwner returns (bool success) {
    	 if (balances[_from] >= _value && balances[_to] + _value >= balances[_to]) {
            balances[_to] += _value;
            balances[_from] -= _value;
            return true;
        } else { return false; }
    }
}