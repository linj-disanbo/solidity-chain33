// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.8.0;

contract User {
    struct Platform {
        bool  exists;
        string pid;
        string endpoint;
    }
    // platform-id to platform info
    mapping(bytes32 => Platform) platforms;
    mapping(bytes32 => mapping(address => bool)) platformManagers;
    // user-id to user in platform
    mapping(bytes32 => mapping (bytes32 => bool)) usersInPlatformMap;
    mapping(bytes32 => bytes32[]) usersInPlatformArray;
    mapping(bytes32 => string) users;

    event LogNewPlatform(
        string pid,
        string endpoint,
        address[] managers
    );
    event LogUpdatePlatformEndPoint(
        string pid,
        string endpoint,
        string oldEndpoint
    );
    event LogDeletePlatform(
        string pid
    );

    event LogAddUserToPlatform(
        string pid,
        string user
    );
    event LogDeleteUserFromPlatform(
        string  pid,
        string user
    );



    modifier isManager(string memory pid) { 
      require(
        platformManagers[keccak256(bytes(pid))][msg.sender] == true,
         "sender is not platform manager"
         );
        _;
    }

    modifier platformExists(string memory pid) {
        require(
            platforms[keccak256(bytes(pid))].exists == true,
            "platforms not exist"
        );
        _;
    }

    modifier userPlatformNotExists(string memory user, string memory pid) {
        require(
            usersInPlatformMap[keccak256(bytes(user))][keccak256(bytes(pid))] == false,
            "platforms   exist"
        );
        _;
    }

            modifier userPlatformExists(string memory user, string memory pid) {
        require(
            usersInPlatformMap[keccak256(bytes(user))][keccak256(bytes(pid))] == true,
            "platforms  not exist"
        );
        _;
    }

    // add platform 
    function addPlatform(string calldata pid, string calldata endpoint, address[] calldata managers) public {
        bytes32 phash = keccak256(bytes(pid));
         require(
        platforms[phash].exists == false,
         "platform exists"
         );
        platforms[phash].exists = true;
        platforms[phash].endpoint = endpoint;
        platforms[phash].pid = pid;
         platformManagers[phash][msg.sender] = true;
         for (uint256 i  = 0; i < managers.length; i = i + 1) {
            platformManagers[phash][managers[i]] = true;
         }

        emit LogNewPlatform(pid, endpoint, managers);
    }

    function deletePlatform(string calldata pid)
     public 
     platformExists(pid)
    {
         platforms[ keccak256(bytes(pid))].exists = false;
        emit LogDeletePlatform(pid);
    }

    // update platform
    function updatePlatform(string calldata pid , string calldata endpoint)  
    public
    isManager(pid)
    {
         bytes32 phash = keccak256(bytes(pid));
        string memory old =  platforms[phash].endpoint;
        platforms[phash].endpoint = endpoint;

        emit LogUpdatePlatformEndPoint(pid, endpoint, old);
    }

    // add user to platform
    function addUserToPlatform(string calldata user, string calldata pid) 
        public
        platformExists(pid)
        isManager(pid)
        userPlatformNotExists(user, pid)
     {
        bytes32 phash = keccak256(bytes(pid));
         bytes32 uhash = keccak256(bytes(user));

         users[uhash]  = user;
         usersInPlatformArray[uhash].push(phash);
         usersInPlatformMap[uhash][phash] = true;
         emit LogAddUserToPlatform(pid, user);
    }

        function deleteUserToPlatform(string calldata user, string calldata pid) 
        public
        platformExists(pid)
        isManager(pid)
        userPlatformExists(user, pid)
         {
        bytes32 phash = keccak256(bytes(pid));
         bytes32 uhash = keccak256(bytes(user));

         usersInPlatformMap[uhash][phash] = false;
        for (uint256 i = 0; i <usersInPlatformArray[uhash].length; i= i+1 ) {
            bytes32 currentPhash = usersInPlatformArray[uhash][i];
            if ( currentPhash != phash ){
                continue;
            }
            if (i < usersInPlatformArray[uhash].length -1) {
                bytes32  last = usersInPlatformArray[uhash][usersInPlatformArray[uhash].length- 1];
                usersInPlatformArray[uhash][i] = last;
            }
            usersInPlatformArray[uhash].pop();
            break;
        }
         emit LogDeleteUserFromPlatform(pid, user);
    }

    // get user info
    function getUser(string memory user) public  view returns(Platform[] memory) {
        bytes32 uhash = keccak256(bytes(user));
        Platform[] memory ps = new Platform[](usersInPlatformArray[uhash].length);
       for ( uint256 i = 0; i < usersInPlatformArray[uhash].length; i=i+1) {
            ps[i] = platforms[usersInPlatformArray[uhash][i]];
       }
       return ps;
    }
}